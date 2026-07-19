package listaccounts

import (
	"bank-account-persistent/events"
	"context"
	"maps"
	"slices"
	"sync"

	"github.com/google/uuid"
	"github.com/terraskye/eventsourcing"
)

type Projector struct {
	mu       sync.RWMutex
	accounts map[uuid.UUID]Account // keyed by AccountID
}

func NewProjector() *Projector {
	// note: should the accounts map be initially populated from event store?
	return &Projector{accounts: map[uuid.UUID]Account{}}
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
