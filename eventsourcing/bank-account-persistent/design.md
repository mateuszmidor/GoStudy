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
event_id (UUID)  |  stream_id (varchar)  |  stream_position (bigint)  |  event_type     |  payload (bytea)  |  metadata (bytea)
uuid-global-1    |  account-uuid-1       |  1                         | AccountCreated  |  JSON bytes       |  (unused)
uuid-global-2    |  account-uuid-1       |  2                         | AccountFunded   |  JSON bytes       |  (unused)
```

Each aggregate (bank account) has its own **stream** — an ordered sequence of events identified by `stream_id`. The `UNIQUE(stream_id, stream_position)` constraint guarantees append-only ordering, preventing concurrent writes to the same stream. The `event_id` column provides a globally unique identifier across all streams (distinct from `stream_position` which is per-stream). The `metadata` column is available but unused in this demo.

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

- **On startup** (`main.go:54-58`): `RebuildFromStore` uses `store.LoadFromAll(ctx, eventsourcing.Any{})` — a different API from `LoadStream` — to read events from **all streams globally** (not just one). Each event is dispatched through the projector's `EventHandlers()`. Events the projector doesn't handle (like `AccountFunded`) produce an `ErrSkippedEvent` which is silently ignored (`projector.go:38-41`). This is the key pattern: **projectors only handle events they care about**.
- **At runtime**: The projector subscribes to the **event bus** (`main.go:59-62`). New `AccountCreated` events are pushed to the projector as they happen, updating the cache.
- **Query handling** (`query.go:30-32`): `GET /accounts` just reads from the in-memory cache — O(1), no DB hit.

**Design decision**: The projector only handles `AccountCreated` — it deliberately ignores `AccountFunded`. This is why the `Account` struct (`query.go:14-19`) has no `Dollars` field (note the comment: `// Dollars uint // for account balance use GetBalance query`). The list view is metadata-only; balance requires a dedicated query via stream replay. This is intentional CQRS separation: different read models serve different purposes.

This is the **CQRS projection pattern**: write-side produces events, read-side maintains a denormalized view.

---

### 5. Event Bus — propagating events to subscribers

The bus (`main.go:51`) uses a **hybrid push+pull** mechanism over PostgreSQL (`schema.sql:20-27`):
- An `AFTER INSERT` trigger on the `events` table fires `pg_notify('eventsourcing_events_inserted', '')` — a lightweight notification that wakes the bus
- The bus then fetches new events from the `events` table, using the `event_subscriptions` table to track each subscriber's last processed position
- Polling interval is 1 second (`main.go:51`: `pgbus.NewEventBus(pool, time.Second)`)
- Events are dispatched to registered handlers (the `listaccounts` projector, in this case)

---

### 6. HTTP routes

| Method | Route | Purpose | Strategy |
|---|---|---|---|
| `POST` | `/accounts` | Create a new account | Command → appends `AccountCreated` event |
| `POST` | `/accounts/{id}/deposits` | Fund an account | Command → appends `AccountFunded` event |
| `GET` | `/accounts` | List all accounts | Query → reads from projector cache |
| `GET` | `/accounts/{id}/balance` | Get account balance | Query → replays event stream |

---

### 7. The complete flow

```
WRITE PATH (commands)                    READ PATH (queries)
══════════════════════                   ═══════════════════

POST /accounts                           GET /accounts/{id}/balance
POST /accounts/{id}/deposits                         │
       │                                        ▼
       ▼                                 ┌────────────────┐
  ┌────────────────┐                     │  Event Store   │
  │ Command Handler │                     │  (PostgreSQL)  │
  │  1. evolve()   │                     │  events table  │
  │     validate    │                     └───────┬────────┘
  │  2. decide()   │                             │
  │     produce     │                             │ replay events → evolve → balance
  └───────┬────────┘                             │
          │ append events                         ▼
          ▼                                 balance (uint)
  ┌────────────────┐
  │  Event Store   │                     GET /accounts
  │  (PostgreSQL)  │                            │
  │  events table  │                            ▼
  └───────┬────────┘                     ┌────────────────┐
          │ pg_notify                    │   Projector    │
          ▼                              │  (in-memory    │
  ┌────────────────┐                     │   cache)       │
  │   Event Bus    │────────────────────►│                │
  │  (poll 1s)     │  dispatch events    └────────────────┘
  └────────────────┘                            │
                                                ▼
                                         accounts list (JSON)
```

**Two distinct read paths**:
- **Projector cache** (`listaccounts`): Fast, eventually consistent. Used for listing accounts.
- **Stream replay** (`getbalance`): Always consistent, reads directly from the event store. Used for balance queries.

---

### Key takeaways for a Go developer

| Concept | How it manifests here |
|---|---|
| **Command-Query separation** | `slices/commands/` and `slices/queries/` are distinct packages with no direct coupling — they share only the `events` package for type definitions |
| **Event sourcing** | State is never stored directly — only events. `accountState` is ephemeral, rebuilt on demand |
| **Projections** | `listaccounts` caches a denormalized view; `getbalance` replays on the fly. Trade-off: staleness vs. consistency |
| **Optimistic concurrency** | The `UNIQUE(stream_id, stream_position)` constraint prevents concurrent appends to the same stream position |
| **Framework** | `github.com/terraskye/eventsourcing` — provides the `CommandHandler`, `EventStore`, `EventBus` abstractions with PostgreSQL implementations |
