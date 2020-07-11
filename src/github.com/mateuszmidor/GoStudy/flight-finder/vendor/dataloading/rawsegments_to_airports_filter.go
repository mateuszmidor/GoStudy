package dataloading

import (
	"airports"
)

type SegmentsToAirportsFilter struct {
}

func NewRawSegmentsToAirportsFilter() *SegmentsToAirportsFilter {
	return &SegmentsToAirportsFilter{}
}

func (f *SegmentsToAirportsFilter) Filter(segments <-chan RawSegment) airports.Airports {
	uniqueCodes := make(map[string]bool)

	for s := range segments {
		uniqueCodes[s.FromAirportCode] = true
		uniqueCodes[s.ToAirportCode] = true
	}

	var ab airports.Builder
	for code := range uniqueCodes {
		ab.Append(code, "")
	}

	return ab.Build()
}
