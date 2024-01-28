package events

// Event is...event :)
type Event int

const (
	// GearUp indicates gear up happened
	GearUp Event = iota

	// GearDown indicates gear down happened
	GearDown

	// DoubleGearDown indicates double gear down happened
	DoubleGearDown
)

// Events is list of events
type Events []Event

// Contains checks if element
func (events Events) Contains(element Event) bool {
	for _, e := range events {
		if e == element {
			return true
		}
	}
	return false
}
