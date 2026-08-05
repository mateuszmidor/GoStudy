package events

// Event mimics eventsourcing.Event from "github.com/terraskye/eventsourcing"
type Event interface {
	AggregateID() string // to correlate event stream with related aggregate
	EventType() string // for event store to identify event type and deserialize event instance
}