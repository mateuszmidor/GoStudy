package getaccount

import (
	"bank-account-persistent/events"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/terraskye/eventsourcing"
)

// GetAccount is the "HandleQuery" input
type GetAccount struct {
	AccountID uuid.UUID
}

// ID is used for query routing if eventsourcing.QueryBus is used
func (g *GetAccount) ID() []byte {
	return []byte ("get-account") // type of the query, not the account id
}

// Account is the "HandleQuery" result
type Account struct {
	AccountID uuid.UUID  
	OwnerName string     
	Balance uint 		 
	CreatedAt time.Time 
}

// accountState is used internally to evolve the account state from events
type accountState struct {
	// AccountID uuid.UUID // not needed, as it comes in the query
	OwnerName string
	Balance   uint
	CreatedAt time.Time
}

// evolve rebuilds state from events
func evolve (state accountState, envelope* eventsourcing.Envelope) accountState {
	switch e:= envelope.Event.(type) {
	case *events.AccountCreated:
		state.CreatedAt = e.CreatedAt
		state.OwnerName = e.OwnerName
		state.Balance = 0
	case *events.AccountFunded:
		state.Balance += e.Dollars
	}

	return state
}

// QueryHandler implements the eventsourcing.QueryHandler interface, and can be used in eventsourcing.QueryBus.
// This is where actual Account is rebuilt from events.
type QueryHandler struct {
	store eventsourcing.EventStore
}

func NewQueryHandler(store eventsourcing.EventStore) *QueryHandler {
	return &QueryHandler{ store: store}
}

// HandleQuery implements the eventsourcing.QueryHandler interface, and can be used in eventsourcing.QueryBus
func (h *QueryHandler) HandleQuery(ctx context.Context, q GetAccount) (Account, error) {
	var noAccount Account // returned in case of error 
	
	// load streams of events for the requested account
	iter, err:= h.store.LoadStream(context.Background(), q.AccountID.String())
	if err != nil {
		if errors.Is(err, eventsourcing.ErrStreamNotFound) {
			return noAccount, fmt.Errorf("get account: account not found %v", q.AccountID)
		}
		return noAccount, fmt.Errorf("get balance: %w", err)
	}

	// rebuild the account state from events
	var state accountState 
	for iter.Next(ctx) {
		state = evolve(state, iter.Value())
	}
	if err := iter.Err(); err != nil {
		return noAccount, err
	}

	// success
	return Account{
		AccountID: q.AccountID,
		OwnerName: state.OwnerName,
		Balance:   state.Balance,
		CreatedAt: state.CreatedAt,
	}, nil
}