package createaccount

import (
	"bank-account/events"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/terraskye/eventsourcing"
)

type CreateAccount struct {
	AccountID uuid.UUID
	OwnerName string
}

func (c CreateAccount) AggregateID() string { return c.AccountID.String() }

type accountState struct {
	exists bool
}

var initialState = func() accountState {
	return accountState{}
}

func evolve(state accountState, envelope *eventsourcing.Envelope) accountState {
	switch e := envelope.Event.(type) {
	case *events.AccountCreated:
		state.exists = true
	default:
		slog.Error(fmt.Sprintf("unknown event %T", e))
	}

	return state
}

func decide(state accountState, cmd CreateAccount) ([]eventsourcing.Event, error) {
	if state.exists {
		return nil, fmt.Errorf("account %v already exists", cmd.AccountID)
	}

	event := &events.AccountCreated{AccountID: cmd.AccountID, OwnerName: cmd.OwnerName, CreatedAt: time.Now()}
	return []eventsourcing.Event{event}, nil
}

func NewHandler(store eventsourcing.EventStore) eventsourcing.CommandHandler[CreateAccount] {
	return eventsourcing.NewCommandHandler(store, initialState, evolve, decide)
}
