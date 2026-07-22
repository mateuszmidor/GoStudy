package events

import "github.com/google/uuid"

type AccountFunded struct {
	AccountID uuid.UUID `json:"account_id"`
	Dollars   uint
}

func (e *AccountFunded) AggregateID() string { return e.AccountID.String() }
func (e *AccountFunded) EventType() string   { return "AccountFunded" }
