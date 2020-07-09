package segment

import (
	"airport"
	"carrier"
)

// ID is segment uid
type ID int

// Segment is smallest part of the journey
type Segment struct {
	from, to airport.AirportID
	carrier  carrier.ID
}

// NewSegment is constructor
func NewSegment(from, to airport.AirportID, carrier carrier.ID) Segment {
	return Segment{from, to, carrier}
}

// From tells where the journey begins
func (s *Segment) From() airport.AirportID {
	return s.from
}

// To tells where the journey ends
func (s *Segment) To() airport.AirportID {
	return s.to
}

// Carrier tells the carrier
func (s *Segment) Carrier() carrier.ID {
	return s.carrier
}
