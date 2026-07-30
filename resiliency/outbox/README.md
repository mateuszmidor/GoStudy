# outbox pattern demo

## Run it

```sh
make rundb
go run .
```

## psql: run SQL interactively

```sh
make rundb
psql -h localhost -p 5432 -U postgres -d postgres
# enter db password: postgres
# then:
select * from outbox; # semicolon means: EXECUTE NOW
```

## Improvements

The current implementation uses **savepoints** to achieve partial progress within a single transaction. When persisting a message fails, the savepoint is rolled back (restoring the transaction to a usable state) and the batch continues. Messages that were successfully persisted are committed; only the failed message is retried on the next poll.

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