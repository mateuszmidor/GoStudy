package application

import (
	"bank-account/events"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Account is the aggregate root for the bank account.
type Account struct {
	ID        uuid.UUID
	OwnerName string
	Balance   uint
	CreatedAt time.Time

	events []events.Event // events are emitted by operations called on Account and used for state storage through event sourcing 
}

// NewAccount creates a new account aggregate and records its creation event.
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
	acc.events = append(acc.events, &events.AccountCreated{
		AccountID: acc.ID,
		OwnerName: acc.OwnerName,
		CreatedAt: acc.CreatedAt,
	})

	return acc, nil
}

// Deposit adds funds to the account and records the funding event.
func (a *Account) Deposit(amount uint) error {
	// check invariants
	var empty uuid.UUID
	if a.ID == empty {
		return errors.New("tried to deposit to uninitialized accout")
	}

	// apply command
	a.Balance += amount

	// emit events to be stored for event sourcing
	a.events = append(a.events, &events.AccountFunded{
		AccountID: a.ID,
		Dollars:   amount,
	})

	return nil
}

// FlushEvents returns pending domain events and clears the local buffer.
func (a *Account) FlushEvents() []events.Event {
	events := a.events
	a.events = nil
	return events
}
