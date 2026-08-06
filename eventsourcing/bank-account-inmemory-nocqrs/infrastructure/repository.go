package infrastructure

import (
	"bank-account/events"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/terraskye/eventsourcing"

	"bank-account/application"
)

var ErrAccountNotFound = errors.New("account not found")
var ErrOptimisticLocking = errors.New("optimistic locking conflict")

// Repository persists and reconstructs accounts from an event store.
type Repository struct {
	store eventsourcing.EventStore
}

// NewRepository builds a repository backed by the provided event store.
func NewRepository(store eventsourcing.EventStore) *Repository {
	return &Repository{
		store: store,
	}
}

// Get rebuilds the state of an account by replaying all events for the given account ID.
func (r *Repository) Get(id uuid.UUID) (*application.Account, error) {
	iter, err := r.store.LoadStream(context.Background(), id.String())
	if err != nil {
		// Keep the public error stable for callers.
		if errors.Is(err, eventsourcing.ErrStreamNotFound) {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}

	// Replay the stream into a single account snapshot.
	var state application.Account
	for iter.Next(context.Background()) {
		state = evolve(state, iter.Value())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}

	state.ID = id
	return &state, nil
}

// GetAll rebuilds the state of all accounts by replaying all events in the event store.
func (r *Repository) List() (application.AccountList, error) {
	iter, err := r.store.LoadFromAll(context.Background(), eventsourcing.Revision(0))
	if err != nil {
		// An empty store is a valid list result.
		if errors.Is(err, eventsourcing.ErrInvalidRevision) {
			return application.AccountList{}, nil
		}
		return nil, err
	}

	// Fold all events into an in-memory account list.
	state := application.AccountList{}
	for iter.Next(context.Background()) {
		state = evolveAll(state, iter.Value())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}

	accounts := make(application.AccountList, 0, len(state))
	for i := range state {
		account := state[i]
		accounts = append(accounts, account)
	}

	return accounts, nil
}

// Save persists the given events to the event store using optimistic concurrency.
// If it returns ErrOptimisticLocking, reload the aggregate, reapply the command,
// and retry Save. Or simply surface the conflict to the caller (HTTP 409).
func (r *Repository) Save(events []events.Event) error {
	if len(events) == 0 {
		return nil
	}

	ctx := context.Background()

	// Repository.Save only accepts one aggregate at a time.
	streamID := events[0].AggregateID()
	for _, event := range events[1:] {
		if event.AggregateID() != streamID {
			return fmt.Errorf("events must belong to the same aggregate")
		}
	}

	// Determine current stream revision so new envelopes get sequential versions.
	revision, err := r.nextStreamRevision(ctx, streamID)
	if err != nil {
		return err
	}

	// Save must receive the expected pre-append stream revision.
	expectedRevision := revision

	// Wrap domain events once before appending them.
	envelopes := make([]eventsourcing.Envelope, len(events))
	for i, event := range events {
		envelopes[i] = eventsourcing.Envelope{
			EventID:    uuid.New(),
			StreamID:   streamID,
			Event:      event,
			Version:    uint64(expectedRevision) + uint64(i),
			OccurredAt: time.Now(),
		}
	}

	_, err = r.store.Save(ctx, envelopes, expectedRevision)
	if err != nil {
		if _, ok := errors.AsType[*eventsourcing.StreamRevisionConflictError](err); ok {
			return fmt.Errorf("%w: %w", ErrOptimisticLocking, err)
		}
	}
	return err
}

// nextStreamRevision returns the current expected revision for appending to streamID.
func (r *Repository) nextStreamRevision(ctx context.Context, streamID string) (eventsourcing.Revision, error) {
	iter, err := r.store.LoadStream(ctx, streamID)
	if err != nil {
		if errors.Is(err, eventsourcing.ErrStreamNotFound) {
			return eventsourcing.Revision(0), nil
		}
		return 0, err
	}

	nextRevision := eventsourcing.Revision(0)
	for iter.Next(ctx) {
		nextRevision++
	}
	if err := iter.Err(); err != nil {
		return 0, err
	}

	return nextRevision, nil
}

func evolve(state application.Account, envelope *eventsourcing.Envelope) application.Account {
	switch e := envelope.Event.(type) {
	case *events.AccountCreated:
		state.CreatedAt = e.CreatedAt
		state.OwnerName = e.OwnerName
		state.Balance = 0
	case *events.AccountFunded:
		state.Balance += e.Dollars
	}

	return state
}

func evolveAll(state application.AccountList, envelope *eventsourcing.Envelope) application.AccountList {
	switch e := envelope.Event.(type) {
	case *events.AccountCreated:
		state = append(state, application.Account{
			ID:        e.AccountID,
			Balance:   0,
			OwnerName: e.OwnerName,
			CreatedAt: e.CreatedAt,
		})
	case *events.AccountFunded:
		// Account always exists - enforced by "FundAccount" command before it emits "AccountFunded" event
		acc := state.GetByID(e.AccountID)
		acc.Balance += e.Dollars
	}

	return state
}
