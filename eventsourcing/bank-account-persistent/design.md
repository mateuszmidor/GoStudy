## Architecture Overview

This is a minimal CQRS/event-sourcing demo. The core insight: **commands produce events, events are the single source of truth, and queries derive state from events**.

---

### 1. Events — the source of truth

There are two event types (`events/account_created.go:9` and `events/account_funded.go:5`):

```go
type AccountCreated struct { AccountID uuid.UUID; OwnerName string; CreatedAt time.Time }
type AccountFunded  struct { AccountID uuid.UUID; Dollars   uint }
```

Events are immutable facts that **already happened**. They implement `AggregateID()` (which stream they belong to) and `EventType()` (for serialization). They are registered in `main.go:30-31` so the framework knows how to deserialize them.

---

### 2. Commands — validation + event production

When `POST /accounts` arrives, `createaccount/command.go` runs a **decide function** (`decide` at line 40):
1. It receives the current **state** (derived by replaying past events on that stream via the `evolve` function at line 29)
2. It validates business rules — e.g. "no duplicate accounts"
3. It returns **new events to append**

The `NewCommandHandler` call (`command.go:51-52`) does the orchestration:
- Loads the event stream from the store (or enforces `NoStream{}` for creation)
- `evolve` replays existing events into `accountState`
- `decide` checks business rules and produces new events
- Framework appends those events atomically to the `events` table in PostgreSQL (using `stream_id + stream_position` uniqueness constraint, see `schema.sql:11`)

`fundaccount/command.go` works identically, but requires `StreamExists{}` instead.

---

### 3. Event Store — durable persistence

PostgreSQL stores events in a single `events` table (`schema.sql:2-12`):

```
stream_id (varchar)   |  stream_position (bigint)  |  event_type  |  payload (bytea)
account-uuid-1        |  1                         | AccountCreated | JSON bytes
account-uuid-1        |  2                         | AccountFunded  | JSON bytes
```

Each aggregate (bank account) has its own **stream** — an ordered sequence of events identified by `stream_id`. The `UNIQUE(stream_id, stream_position)` constraint guarantees append-only ordering, preventing concurrent writes to the same stream.

---

### 4. Queries — two strategies for deriving state

**Strategy A — Event stream replay** (`getbalance/query.go:41-58`)

On each `GET /accounts/{id}/balance` request, `HandleQuery`:
1. Loads the entire event stream for that account from the store via `store.LoadStream`
2. Replays every event through the `evolve` function: `AccountFunded` adds `Dollars` to the state
3. Returns the accumulated balance

This is the purest form — **no cached state, always consistent**.

**Strategy B — Projector / cache** (`listaccounts/`)

The `Projector` maintains an in-memory `map[uuid.UUID]Account`:

- **On startup** (`main.go:54-58`): `RebuildFromStore` replays **all** streams from the store, calling each event through the projector's `OnAccountCreated` handler. This populates the cache to catch up with everything that happened before this process was alive.
- **At runtime**: The projector subscribes to the **event bus** (`main.go:59-62`). New `AccountCreated` events are pushed to the projector as they happen, updating the cache.
- **Query handling** (`query.go:30-32`): `GET /accounts` just reads from the in-memory cache — O(1), no DB hit.

This is the **CQRS projection pattern**: write-side produces events, read-side maintains a denormalized view.

---

### 5. Event Bus — propagating events to subscribers

The bus (`main.go:51`) uses **PostgreSQL LISTEN/NOTIFY** (`schema.sql:20-27`):
- An `AFTER INSERT` trigger on the `events` table fires `pg_notify('eventsourcing_events_inserted', '')`
- The bus polls for new events via `event_subscriptions` table tracking the last processed position per subscriber
- Dispatches events to registered handlers (the `listaccounts` projector, in this case)

---

### 6. The complete flow

```
HTTP POST /accounts                 HTTP GET /accounts                 HTTP GET /accounts/{id}/balance
       │                                  │                                    │
       ▼                                  │                                    │
  Command handler                         │                                    │
  ┌──────────────┐                        │                                    │
  │ 1. evolve()  │◄─── load stream ───────┤                                    │
  │    (validate)│                        │                                    │
  │ 2. decide()  │                        │                                    │
  │    (produce) │                        │                                    │
  └──────┬───────┘                        │                                    │
         │ append events                  │                                    │
         ▼                                │                                    │
  ┌──────────────┐                        │                                    │
  │  Event Store │                        │                                    │
  │  (PostgreSQL)│                        │                                    │
  │  events table│                        │                                    │
  └──────┬───────┘                        │                                    │
         │ pg_notify                      │                                    │
         ▼                                │                                    │
  ┌──────────────┐     on next poll       │                                    │
  │  Event Bus   │ ───────────────────►   │    ┌──────────────────────────┐    │
  │  (LISTEN/    │                        ├───►│  Projector (cache)       │    │
  │   NOTIFY)    │                        │    │  reads: O(1) map lookup  │    │
  └──────────────┘                        │    └──────────────────────────┘    │
                                          │                                    │
                                          │    ┌──────────────────────────┐    │
                                          └───►│  Stream Replay            │◄───
                                               │  evolves events → state  │
                                               └──────────────────────────┘
```

---

### Key takeaways for a Go developer

| Concept | How it manifests here |
|---|---|
| **Command-Query separation** | `slices/commands/` and `slices/queries/` are distinct packages with zero shared code |
| **Event sourcing** | State is never stored directly — only events. `accountState` is ephemeral, rebuilt on demand |
| **Projections** | `listaccounts` caches a denormalized view; `getbalance` replays on the fly. Trade-off: staleness vs. consistency |
| **Idempotency** | Commands use `stream_position + UNIQUE` constraint to prevent duplicate appends |
| **Framework** | `github.com/terraskye/eventsourcing` — provides the `CommandHandler`, `EventStore`, `EventBus` abstractions with PostgreSQL implementations |
