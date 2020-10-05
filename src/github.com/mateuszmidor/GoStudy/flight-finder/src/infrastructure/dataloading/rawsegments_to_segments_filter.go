package dataloading

import (
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/airports"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/carriers"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/segments"
)

type RawSegmentsToSegmentsFilter struct {
	airports airports.Airports
	carriers carriers.Carriers
}

func NewRawSegmentsToSegmentsFilter(airports airports.Airports, carriers carriers.Carriers) *RawSegmentsToSegmentsFilter {
	return &RawSegmentsToSegmentsFilter{airports, carriers}
}

func (f *RawSegmentsToSegmentsFilter) Filter(segs <-chan RawSegment) segments.Segments {
	sb := segments.NewBuilder(f.airports, f.carriers)

	for s := range segs {
		sb.Append(s.FromAirportCode, s.ToAirportCode, s.CarrierCode)
	}

	return sb.Build()
}
