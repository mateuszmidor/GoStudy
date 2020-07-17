package asjson

import (
	"airports"
	"carriers"
)

// SegmentView is json view of segments.Segment
type SegmentView struct {
	Carrier   *CarrierView `json:"carrier"`
	ToAirport *AirportView `json:"to_airport"`
}

// NewJSONSegmentView is constructor
func NewJSONSegmentView(c *carriers.Carrier, a *airports.Airport) *SegmentView {
	return &SegmentView{
		Carrier:   NewJSONCarrierView(c),
		ToAirport: NewJSONAirportView(a),
	}
}
