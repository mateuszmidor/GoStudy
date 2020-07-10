package dataloading

import (
	"airport"
	"carrier"
	"segment"
)

type RawSegmentsToSegmentsFilter struct {
	airports airport.Airports
	carriers carrier.Carriers
}

func NewRawSegmentsToSegmentsFilter(airports airport.Airports, carriers carrier.Carriers) *RawSegmentsToSegmentsFilter {
	return &RawSegmentsToSegmentsFilter{airports, carriers}
}

func (f *RawSegmentsToSegmentsFilter) Filter(segments <-chan segment.RawSegment) segment.Segments {
	sb := segment.NewBuilder(f.airports, f.carriers)

	for s := range segments {
		sb.Append(s.FromAirportCode, s.ToAirportCode, s.CarrierCode)
	}

	return sb.Build()
}
