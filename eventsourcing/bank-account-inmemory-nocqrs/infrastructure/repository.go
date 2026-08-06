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
	var account application.Account
	for iter.Next(context.Background()) {
		account = evolve(account, iter.Value())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}

	return &account, nil
}

// List rebuilds the state of all accounts by replaying all events in the event store.
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
	list := application.AccountList{}
	for iter.Next(context.Background()) {
		list = evolveList(list, iter.Value())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}

	return list, nil
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

// evolveList folds a single event into the current state of all accounts.
func evolveList(state application.AccountList, envelope *eventsourcing.Envelope) application.AccountList {
	id := uuid.MustParse(envelope.Event.AggregateID()) // will always parse; the ID is a stringified UUID
	account := state.GetOrInsert(id) // get account, or insert a new one if it doesn't exist yet
	*account = evolve(*account, envelope) // apply the event to the account
	return state // return the updated list
}

// evolve folds a single event into the current state of an account.
func evolve(state application.Account, envelope *eventsourcing.Envelope) application.Account {
	return apply(state, envelope.Event) // simply apply the event to the account state
}

// apply folds a single event into the current state of an account.
// The valid event order is guaranteed by the repository, so we can safely apply events in the order they are received.
func apply (state application.Account, event events.Event) application.Account {
	switch e := event.(type) {
	case *events.AccountCreated:
		state.ID = e.AccountID
		state.CreatedAt = e.CreatedAt
		state.OwnerName = e.OwnerName
		state.Balance = 0
	case *events.AccountFunded:
		state.Balance += e.Dollars
	}

	return state
}