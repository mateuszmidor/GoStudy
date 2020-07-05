package segment

import "airport"

// ID is segment uid
type ID int

// Segment is smallest part of the journey
type Segment struct {
	from, to airport.AirportID
}

// NewSegment is constructor
func NewSegment(from, to airport.AirportID) Segment {
	return Segment{from, to}
}

// From tells where the journey begins
func (s *Segment) From() airport.AirportID {
	return s.from
}

// To tells where the journey ends
func (s *Segment) To() airport.AirportID {
	return s.to
}
