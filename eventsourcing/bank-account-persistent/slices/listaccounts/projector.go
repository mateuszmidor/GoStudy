package listaccounts

import (
	"bank-account-persistent/events"
	"context"
	"errors"
	"fmt"
	"maps"
	"slices"
	"sync"

	"github.com/google/uuid"
	"github.com/terraskye/eventsourcing"
)

// Projector caches the accounts
type Projector struct {
	mu       sync.RWMutex
	accounts map[uuid.UUID]Account // keyed by AccountID
}

func NewProjector() *Projector {
	return &Projector{accounts: map[uuid.UUID]Account{}}
}

// RebuildFromStore populates the projector cache by replaying all events from the event store.
// Call this before subscribing to the event bus to ensure the projector is up to date.
func (p *Projector) RebuildFromStore(ctx context.Context, store eventsourcing.EventStore) error {
	iter, err := store.LoadFromAll(ctx, eventsourcing.Any{})
	if err != nil {
		return fmt.Errorf("rebuild from store: %w", err)
	}

	handleEvent := p.EventHandlers().Handle
	for iter.Next(ctx) {
		event := iter.Value().Event
		if err := handleEvent(ctx, event); err != nil {
			var skipped *eventsourcing.ErrSkippedEvent
			if !errors.As(err, &skipped) {
				return fmt.Errorf("handle event %s: %w", event.EventType(), err)
			}
		}
	}

	return iter.Err()
}

func (p *Projector) GetAll() []Account {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return slices.Collect(maps.Values(p.accounts))
}

// OnAccountCreated is called by the event bus when a AccountCreated event arrives.
func (p *Projector) OnAccountCreated(_ context.Context, e *events.AccountCreated) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.accounts[e.AccountID] = Account{
		ID:        e.AccountID,
		OwnerName: e.OwnerName,
		CreatedAt: e.CreatedAt,
	}

	return nil
}

// handle AccountDeleted event when added
// func (p *Projector) OnAccountDeleted(_ context.Context, e *events.AccountDeleted) error {}

// EventHandlers returns the typed handlers to register on the event bus.
func (p *Projector) EventHandlers() *eventsourcing.EventGroupProcessor {
	return eventsourcing.NewEventGroupProcessor(
		eventsourcing.OnEvent(p.OnAccountCreated),
	)
}
