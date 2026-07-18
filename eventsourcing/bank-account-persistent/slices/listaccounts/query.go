package listaccounts

import (
	"bank-account-persistent/events"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/terraskye/eventsourcing"
)

type ListAccounts struct{}

func (q ListAccounts) ID() []byte { return []byte("list-accounts") }

type Account struct {
	ID        string
	Dollars   uint
	OwnerName string
	CreatedAt time.Time
}

type AccountList struct {
	Accounts []Account
}

func (l *AccountList) getByID(accountID string) *Account {
	for i := range l.Accounts {
		if l.Accounts[i].ID == accountID {
			return &l.Accounts[i]
		}
	}
	return nil
}

func evolve(state *AccountList, envelope *eventsourcing.Envelope) *AccountList {
	switch e := envelope.Event.(type) {
	case *events.AccountCreated:
		state.Accounts = append(state.Accounts, Account{
			ID:        e.AccountID.String(),
			Dollars:   0,
			OwnerName: e.OwnerName,
			CreatedAt: e.CreatedAt,
		})
	case *events.AccountFunded:
		// Account always exists - enforced by "FundAccount" command before it emits "AccountFunded" event
		acc := state.getByID(e.AccountID.String())
		acc.Dollars += e.Dollars
	}

	return state
}

type QueryHandler struct {
	store eventsourcing.EventStore
}

func NewQueryHandler(store eventsourcing.EventStore) *QueryHandler {
	return &QueryHandler{store: store}
}

func (h *QueryHandler) HandleQuery(ctx context.Context, _ ListAccounts) (*AccountList, error) {
	iter, err := h.store.LoadFromAll(ctx, eventsourcing.Revision(0))
	if err != nil {
		if errors.Is(err, eventsourcing.ErrInvalidRevision) {
			return &AccountList{Accounts: make([]Account, 0)}, nil
		}
		return nil, fmt.Errorf("list accounts: %w", err)
	}

	accounts := &AccountList{Accounts: make([]Account, 0)}
	for iter.Next(ctx) {
		accounts = evolve(accounts, iter.Value())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}
