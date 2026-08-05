package application

import (
	"bank-account/events"
	"errors"
	"time"

	"github.com/google/uuid"
)

type AccountList []Account

func (list AccountList) GetByID(accountID uuid.UUID) *Account {
	for account := range list {
		if list[account].ID == accountID {
			return &list[account]
		}
	}
	return nil
}

type Account struct {
	ID        uuid.UUID
	OwnerName string
	Balance   uint
	CreatedAt time.Time

	Events []events.Event
}

func NewAccount(ownerName string) (*Account, error) {
	// check invariants
	if ownerName == "" {
		return nil, errors.New("owner name cannot be empty")
	}

	// apply command
	acc := &Account{
		ID:        uuid.New(),
		OwnerName: ownerName,
		Balance:   0,
		CreatedAt: time.Now(),
	}

	// emit events to be stored for event sourcing
	acc.Events = append(acc.Events, &events.AccountCreated{
		AccountID: acc.ID,
		OwnerName: acc.OwnerName,
		CreatedAt: acc.CreatedAt,
	})

	return acc, nil
}

func (a *Account) Deposit(amount uint) error {
	// check invariants
	var empty uuid.UUID
	if a.ID == empty {
		return errors.New("tried to deposit to uninitialized accout")
	}

	// apply command
	a.Balance += amount

	// emit events to be stored for event sourcing
	a.Events = append(a.Events, &events.AccountFunded{
		AccountID: a.ID,
		Dollars:   amount,
	})

	return nil
}

func (a *Account) FlushEvents() []events.Event {
	events := a.Events
	a.Events = nil
	return events
}
