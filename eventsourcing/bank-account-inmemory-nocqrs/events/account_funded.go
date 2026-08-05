package events

import "github.com/google/uuid"

// AccountFunded implements Event and records a deposit made into an account.
type AccountFunded struct {
	AccountID uuid.UUID `json:"account_id"`
	Dollars   uint `json:"dollars"`
}

// AggregateID returns the account ID this event belongs to.
func (e *AccountFunded) AggregateID() string { return e.AccountID.String() }

// EventType identifies the event for event store serialization.
func (e *AccountFunded) EventType() string { return "AccountFunded" }
