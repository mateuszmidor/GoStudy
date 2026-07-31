# outbox pattern demo

## Run it

```sh
make rundb # runs postgres with "outbox" table
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

## How it works

```
┌─────────────────────────────────────────────────────────────────────────────────────┐
│                        WRITE SIDE  (single DB transaction)                          │
└─────────────────────────────────────────────────────────────────────────────────────┘

   main.go
     │
     ▼
 ┌──────────┐   Publisher   ┌──────────────────────────────────┐
 │ *sql.DB  │───Publish()──▶│ RunInTransaction                 │
 └──────────┘               │                                  │
                            │  BEGIN                           │
                            │    ▼                             │
                            │  PersistMessage (INSERT INTO     │
                            │  outbox ... ON CONFLICT UPDATE)  │
                            │    ▼                             │
                            │  tx.Commit                       │
                            └──────────────────────────────────┘
                                          │
                                outbox row is now in DB,
                                waiting for relay to pick up


┌─────────────────────────────────────────────────────────────────────────────────────┐
│                    READ SIDE  (relay polling loop, every 3s)                        │
└─────────────────────────────────────────────────────────────────────────────────────┘

   relay.Run (OutboxRelay)
     │
     ▼
 ┌──────────────────┐
 │ time.After(3s)   │◀──── wait for next tick
 └────────┬─────────┘
          ▼
 ┌──────────────────┐  BEGIN   ┌──────────────────────────────────┐
 │ *sql.DB          │─────────▶│ *sql.Tx (long-lived for batch)   │
 └──────────────────┘          └──────────────────────────────────┘
                                        │
                                        ▼
                         ┌──────────────────────────────┐
                         │ SELECT id, message_name, ... │
                         │  FROM outbox                 │
                         │  WHERE processed_at IS NULL  │
                         │    AND failed != true        │
                         │  ORDER BY occurred_at        │
                         │  LIMIT 100                   │
                         │  FOR UPDATE SKIP LOCKED      │  ◄── row locks; skip rows
                         └──────────────────────────────┘      locked by other workers
                                        │
                                        ▼ messages[]
                         ┌──────────────────────────────┐
                         │  for msg in messages:        │
                         │                              │
                         │  ┌────────────────────────┐  │
                         │  │ visitorFunc(ctx, msg)  │  │  ◄── external I/O happens
                         │  │                        │  │      FIRST, before savepoint
                         │  │  publish message       │  │
                         │  │    │                   │  │
                         │  │    ├── ok ──▶          │  │
                         │  │      MarkAsProcessed() │  │
                         │  │       (in-memory only) │  │
                         │  │    │                   │  │
                         │  │    └── err ▶           │  │
                         │  │      MarkAsFailed      │  │
                         │  │       (failed=true,    │  │
                         │  │        failCount++,    │  │
                         │  │        in-memory only) │  │
                         │  └───────────┬────────────┘  │
                         │              ▼               │
                         │  ┌────────────────────────┐  │
                         │  │ SAVEPOINT persist_msg  │  │  ◄── protects DB state only
                         │  └───────────┬────────────┘  │
                         │              ▼               │
                         │  ┌────────────────────────┐  │
                         │  │ PersistMessage (UPSERT)│  │  ◄── write in-memory state
                         │  │  processed_at /        │  │      (processed_at, fail_count,
                         │  │  fail_count /          │  │       failed, failure_reason)
                         │  │  failed / etc.         │  │      to DB
                         │  └───────────┬────────────┘  │
                         │              │               │
                         │        ┌─────┴─────┐         │
                         │        ▼           ▼         │
                         │    success       error       │
                         │        │           │         │
                         │        ▼           ▼         │
                         │  RELEASE       ROLLBACK TO   │
                         │  SAVEPOINT     SAVEPOINT     │  ◄── restore tx from Postgres
                         │                persist_msg   │      "aborted" state so
                         │        │           │         │      remaining messages can
                         │        └─────┬─────┘         │      still be processed
                         │              ▼               │
                         │         next message         │
                         └──────────────────────────────┘
                                        │
                                        ▼
                                  ┌───────────┐
                                  │ tx.Commit │  ◄── all successful persists are
                                  └───────────┘      saved atomically


┌─────────────────────────────────────────────────────────────────────────────────────┐
│                            MESSAGE LIFECYCLE                                        │
└─────────────────────────────────────────────────────────────────────────────────────┘

   publish OK                              publish FAIL
      │                                        │
      ▼                                        ▼
 MarkAsProcessed()                      MarkAsFailed(reason)
   processedAt = now()                    failed = true
                                          failCount++
                                          failureReason = reason
      │                                        │
      ▼                                        ▼
 PersistMessage → DB                    PersistMessage → DB
   processed_at = now()                   failed = true
                                          fail_count = 1
      │                                        │
      ▼                                        ▼
 SELECT ... WHERE                             SELECT ... WHERE
   processed_at IS NULL                         processed_at IS NULL
   AND failed != true                           AND failed != true
      │                                        │
      ▼                                        ▼
 row EXCLUDED from polls                   row EXCLUDED from polls
 (will never be retried)                   (will never be retried)


    Messages that fail are retried with exponential backoff
    (`power(2, fail_count-1)` second delay — first retry waits 1s, then 2s, 4s, etc.). The `WithMaxAttempts(N)` option controls
    how many total processing attempts are allowed before
    the message is permanently marked as failed.
```