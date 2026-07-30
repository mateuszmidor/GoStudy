# outbox

This package implements realiable outbox pattern.

## Improvements

The current implementation uses **savepoints** to achieve partial progress within a single transaction. When persisting a handled message fails in publisher, the savepoint is rolled back (restoring the transaction to a usable state) and the batch continues. Messages that were successfully persisted are committed; only the failed message is retried on the next poll.

### Drawbacks of the current approach

- **Long-held transaction with row locks**: The entire batch (up to 100 messages) is processed within one transaction that holds `FOR UPDATE` row locks. If the publish operation (external I/O) is slow, the transaction remains open for the duration, which can cause connection pool pressure and lock contention at scale.
- **Extra round trips**: Each message incurs two additional statements (`SAVEPOINT` + `RELEASE`), adding overhead for large batches.

### Production-ready alternative: Claim Pattern

For high-throughput production systems, the recommended approach is the **claim pattern** (also known as leasing). The key idea: **never hold a database transaction open while doing external I/O**.

**How it works:**

1. **Claim phase** (short transaction): Select a batch of unprocessed messages with `FOR UPDATE SKIP LOCKED`, mark them as claimed (e.g., set `claimed_at = now()`), and commit immediately. This releases the row locks.

2. **Process phase** (per-message, short transactions): For each claimed message, run the handler (publish to broker) and update the message state (`processed_at`, clear `claimed_at`) in its own short transaction. If processing fails, the message's `claimed_at` remains set.

3. **Cleanup** (background job): Periodically reclaim messages whose `claimed_at` is older than a timeout (e.g., 5 minutes) — these are messages whose worker crashed or got stuck.

**Benefits:**
- Short transactions throughout (no long lock hold)
- Failures are fully isolated (one message's failure doesn't affect others)
- Scales well horizontally (multiple workers, no lock contention after claim)
- Stale claims are recoverable (no manual intervention needed)

**Schema changes required:**
```sql
ALTER TABLE outbox ADD COLUMN claimed_at timestamp with time zone;
ALTER TABLE outbox ADD COLUMN claimed_by text; -- optional, for multi-worker safety

CREATE INDEX idx_outbox_claimable ON outbox (occurred_at)
WHERE processed_at IS NULL AND failed != true AND claimed_at IS NULL;
```

## Critical Improvements for Production

### 0. Claim pattern

Described above

### 1. Exponential Backoff ✅ implemented

Delay based on `fail_count` to avoid overwhelming downstream systems during outages.

**Approach:**
- Calculate delay as `power(2, fail_count-1)` seconds (e.g., 1s, 2s, 4s, 8s, 16s...)
- Modify polling query to only fetch messages where `occurred_at + delay < now()`

**Query modification:**
```sql
SELECT * FROM outbox
WHERE processed_at IS NULL 
  AND failed != true
  AND occurred_at + (power(2, fail_count - 1) * interval '1 second') < now()
ORDER BY occurred_at
LIMIT $1
FOR UPDATE SKIP LOCKED
```

### 2. Max Attempts & Dead Letter Queue

Move permanently failed messages to a separate table after N attempts to prevent infinite retries.

**Approach:**
- Define max attempts threshold (e.g., 10 attempts)
- When `fail_count + 1 >= maxAttempts`, mark as `failed = true` (dead letter)
- Background job or manual process handles dead letter messages

**Schema changes:**
```sql
CREATE TABLE dead_letter (
    id             uuid                  not null,
    event_name     text                  not null,
    event_data     jsonb                 not null,
    occurred_at    timestamp with time zone not null,
    fail_count     integer               not null,
    failure_reason text,
    failed_at      timestamp with time zone default now() not null,
    constraint dead_letter_pk primary key (id)
);
```

**Implementation:**
```go
if msg.FailCount()+1 >= maxAttempts {
    msg.MarkAsFailed(reason)
    // Move to dead_letter table
    // Delete from outbox table
}
```

### 3. Idempotency Support

Prevent duplicate message processing when messages are delivered multiple times.

**Approach:**
- Add `idempotency_key` column (unique business identifier)
- Check for duplicate keys before processing
- Store processed idempotency keys with TTL

**Schema changes:**
```sql
ALTER TABLE outbox ADD COLUMN idempotency_key text;
CREATE UNIQUE INDEX idx_outbox_idempotency ON outbox (idempotency_key) 
    WHERE idempotency_key IS NOT NULL AND processed_at IS NULL;

-- Optional: track processed keys for deduplication
CREATE TABLE processed_keys (
    idempotency_key text primary key,
    processed_at    timestamp with time zone default now() not null
);
```

**Implementation:**
```go
// Before processing, check if already processed
if alreadyProcessed(msg.IdempotencyKey()) {
    msg.MarkAsProcessed()
    return nil
}
```
