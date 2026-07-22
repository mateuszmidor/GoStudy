package events

import (
	"time"

	"github.com/google/uuid"
)

type AccountCreated struct {
	AccountID uuid.UUID `json:"account_id"`
	OwnerName string    `json:"ownerr_name"`
	CreatedAt time.Time `json:"created_at"`
}

func (e *AccountCreated) AggregateID() string { return e.AccountID.String() }
func (e *AccountCreated) EventType() string   { return "AccountCreated" }
