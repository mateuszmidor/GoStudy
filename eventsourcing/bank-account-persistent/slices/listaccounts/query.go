package listaccounts

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ListAccounts struct{}

func (q ListAccounts) ID() []byte { return []byte("list-accounts") }

type Account struct {
	ID        uuid.UUID
	OwnerName string
	CreatedAt time.Time
}

// type AccountList struct {
// 	Accounts []Account
// }

// func (l *AccountList) getByID(accountID string) *Account {
// 	for i := range l.Accounts {
// 		if l.Accounts[i].ID == accountID {
// 			return &l.Accounts[i]
// 		}
// 	}
// 	return nil
// }

// func evolve(state *AccountList, envelope *eventsourcing.Envelope) *AccountList {
// 	switch e := envelope.Event.(type) {
// 	case *events.AccountCreated:
// 		state.Accounts = append(state.Accounts, Account{
// 			ID:        e.AccountID.String(),
// 			Dollars:   0,
// 			OwnerName: e.OwnerName,
// 			CreatedAt: e.CreatedAt,
// 		})
// 	case *events.AccountFunded:
// 		// Account always exists - enforced by "FundAccount" command before it emits "AccountFunded" event
// 		acc := state.getByID(e.AccountID.String())
// 		acc.Dollars += e.Dollars
// 	}

// 	return state
// }

type QueryHandler struct {
	projector *Projector
}

func NewQueryHandler(projector *Projector) *QueryHandler {
	return &QueryHandler{projector: projector}
}

// HandleQuery returns the cached account list
func (h *QueryHandler) HandleQuery(ctx context.Context, _ ListAccounts) ([]Account, error) {
	// iter, err := h.store.LoadFromAll(ctx, eventsourcing.Any{})
	// if err != nil {
	// 	if errors.Is(err, eventsourcing.ErrInvalidRevision) {
	// 		return &AccountList{Accounts: make([]Account, 0)}, nil
	// 	}
	// 	return nil, fmt.Errorf("list accounts: %w", err)
	// }

	// accounts := &AccountList{Accounts: make([]Account, 0)}
	// for iter.Next(ctx) {
	// 	accounts = evolve(accounts, iter.Value())
	// }
	// if err := iter.Err(); err != nil {
	// 	return nil, err
	// }

	// return accounts, nil

	// return accounts from cache
	return h.projector.GetAll(), nil
}
