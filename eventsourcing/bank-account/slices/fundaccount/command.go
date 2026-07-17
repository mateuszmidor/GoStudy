package fundaccount

import (
	"bank-account/events"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/terraskye/eventsourcing"
)

type FundAccount struct {
	AccountID uuid.UUID
	Dollars   uint
}

func (f FundAccount) AggregateID() string { return f.AccountID.String() }

type accountState struct {
	dollars uint
}

func initialState() accountState {
	return accountState{}
}

func evolve(state accountState, envelope *eventsourcing.Envelope) accountState {
	switch e := envelope.Event.(type) {
	case *events.AccountCreated:
		{
		}
	case *events.AccountFunded:
		state.dollars += e.Dollars
	default:
		slog.Error(fmt.Sprintf("unknown event %T", e))
	}

	return state
}

func decide(state accountState, cmd FundAccount) ([]eventsourcing.Event, error) {
	event := &events.AccountFunded{AccountID: cmd.AccountID, Dollars: cmd.Dollars}
	return []eventsourcing.Event{event}, nil
}

func NewHandler(store eventsourcing.EventStore) eventsourcing.CommandHandler[FundAccount] {
	return eventsourcing.NewCommandHandler(store, initialState, evolve, decide)
}
