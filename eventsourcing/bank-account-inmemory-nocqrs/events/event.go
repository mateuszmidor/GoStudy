package events

// Event mimics eventsourcing.Event from "github.com/terraskye/eventsourcing"
type Event interface {
	AggregateID() string
	EventType() string
}