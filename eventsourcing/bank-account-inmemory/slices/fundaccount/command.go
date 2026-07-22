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

// AggregateID is used by the eventsourcing framework to load correct event stream of events for "evolve" function
func (f FundAccount) AggregateID() string { return f.AccountID.String() }

type accountState struct {
	created bool
}

func initialState() accountState {
	return accountState{}
}

func evolve(state accountState, envelope *eventsourcing.Envelope) accountState {
	switch e := envelope.Event.(type) {
	case *events.AccountCreated:
		state.created = true
	case *events.AccountFunded:
		break // if it was funded before that's ok
	default:
		slog.Error(fmt.Sprintf("unknown event %T", e))
	}

	return state
}

func decide(state accountState, cmd FundAccount) ([]eventsourcing.Event, error) {
	// account must be created before being funded so check it here
	// this actually is a double-check after the eventsourcing.StreamExists{} requirement
	if !state.created {
		return nil, fmt.Errorf("account id %v not created yet", cmd.AccountID)
	}
	event := &events.AccountFunded{AccountID: cmd.AccountID, Dollars: cmd.Dollars}
	return []eventsourcing.Event{event}, nil
}

func NewHandler(store eventsourcing.EventStore) eventsourcing.CommandHandler[FundAccount] {
	return eventsourcing.NewCommandHandler(store, initialState, evolve, decide, eventsourcing.WithStreamState(eventsourcing.StreamExists{})) // aggregate with such ID must already exist
}
