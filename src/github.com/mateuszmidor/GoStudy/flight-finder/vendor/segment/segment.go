package segment

import (
	"airport"
	"carrier"
)

// ID is segment uid
type ID int

// NullID means no such segment
const NullID = ID(-1)

// Segment is smallest part of the journey
type Segment struct {
	from    airport.ID
	to      airport.ID
	carrier carrier.ID
}

// NewSegment is constructor
func NewSegment(from, to airport.ID, carrier carrier.ID) Segment {
	return Segment{from, to, carrier}
}

// From tells where the journey begins
func (s *Segment) From() airport.ID {
	return s.from
}

// To tells where the journey ends
func (s *Segment) To() airport.ID {
	return s.to
}

// Carrier tells the carrier
func (s *Segment) Carrier() carrier.ID {
	return s.carrier
}
