package events

import (
	"time"

	"github.com/google/uuid"
)

// AccountCreated implements Event and records the creation of a new account.
type AccountCreated struct {
	AccountID uuid.UUID `json:"account_id"`
	OwnerName string    `json:"owner_name"`
	CreatedAt time.Time `json:"created_at"`
}

// AggregateID returns the account ID this event belongs to.
func (e *AccountCreated) AggregateID() string { return e.AccountID.String() }

// EventType identifies the event for event store serialization.
func (e *AccountCreated) EventType() string { return "AccountCreated" }
