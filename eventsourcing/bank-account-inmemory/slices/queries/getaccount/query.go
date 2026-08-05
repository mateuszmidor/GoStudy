package getaccount

import (
	"bank-account/events"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/terraskye/eventsourcing"
)

type GetAccount struct {
	AccountID uuid.UUID
}

func (g *GetAccount) ID() []byte {
	return []byte("get-account")
}

type Account struct {
	AccountID uuid.UUID
	OwnerName string
	Balance   uint
	CreatedAt time.Time
}

type accountState struct {
	OwnerName string
	Balance   uint
	CreatedAt time.Time
}

func evolve(state accountState, envelope *eventsourcing.Envelope) accountState {
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

type QueryHandler struct {
	store eventsourcing.EventStore
}

func NewQueryHandler(store eventsourcing.EventStore) *QueryHandler {
	return &QueryHandler{store: store}
}

func (h *QueryHandler) HandleQuery(ctx context.Context, q GetAccount) (Account, error) {
	var noAccount Account

	iter, err := h.store.LoadStream(ctx, q.AccountID.String())
	if err != nil {
		if errors.Is(err, eventsourcing.ErrStreamNotFound) {
			return noAccount, fmt.Errorf("get account: account not found %v", q.AccountID)
		}
		return noAccount, fmt.Errorf("get account: %w", err)
	}

	var state accountState
	for iter.Next(ctx) {
		state = evolve(state, iter.Value())
	}
	if err := iter.Err(); err != nil {
		return noAccount, err
	}

	return Account{
		AccountID: q.AccountID,
		OwnerName: state.OwnerName,
		Balance:   state.Balance,
		CreatedAt: state.CreatedAt,
	}, nil
}
