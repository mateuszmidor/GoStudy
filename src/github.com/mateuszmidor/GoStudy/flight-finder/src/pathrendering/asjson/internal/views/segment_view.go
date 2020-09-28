package views

import (
	"github.com/mateuszmidor/GoStudy/flight-finder/src/airports"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/carriers"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/nations"
)

// Segment is json view of segments.Segment
type Segment struct {
	Carrier   *Carrier `json:"carrier"`
	ToAirport *Airport `json:"to_airport"`
}

// NewJSONSegmentView is constructor
func NewJSONSegmentView(c *carriers.Carrier, a *airports.Airport, n *nations.Nation) *Segment {
	return &Segment{
		Carrier:   NewJSONCarrierView(c),
		ToAirport: NewJSONAirportView(a, n),
	}
}
