package views

import (
	"airports"
	"carriers"
)

// Segment is json view of segments.Segment
type Segment struct {
	Carrier   *Carrier `json:"carrier"`
	ToAirport *Airport `json:"to_airport"`
}

// NewJSONSegmentView is constructor
func NewJSONSegmentView(c *carriers.Carrier, a *airports.Airport) *Segment {
	return &Segment{
		Carrier:   NewJSONCarrierView(c),
		ToAirport: NewJSONAirportView(a),
	}
}