package getbalance

import (
	"bank-account-persistent/events"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/terraskye/eventsourcing"
)

type GetBalance struct {
	AccountID uuid.UUID
}

func (q GetBalance) ID() []byte { return []byte("get-balance") }

type accountState struct {
	Dollars uint
}

func evolve(state accountState, envelope *eventsourcing.Envelope) accountState {
	switch e := envelope.Event.(type) {
	case *events.AccountFunded:
		state.Dollars += e.Dollars
	}

	return state
}

type QueryHandler struct {
	store eventsourcing.EventStore
}

func NewQueryHandler(store eventsourcing.EventStore) *QueryHandler {
	return &QueryHandler{store: store}
}

// HandleQuery rebuilds the account balance event by event
func (h *QueryHandler) HandleQuery(ctx context.Context, q GetBalance) (uint, error) {
	iter, err := h.store.LoadStream(ctx, q.AccountID.String())
	if err != nil {
		if errors.Is(err, eventsourcing.ErrInvalidRevision) {
			return 0, nil
		}
		return 0, fmt.Errorf("get balance: %w", err)
	}

	var state accountState
	for iter.Next(ctx) {
		state = evolve(state, iter.Value())
	}
	if err := iter.Err(); err != nil {
		return 0, err
	}

	return state.Dollars, nil
}
