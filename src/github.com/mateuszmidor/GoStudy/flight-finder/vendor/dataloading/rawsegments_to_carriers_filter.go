package dataloading

import (
	"carrier"
)

type SegmentsToCarriersFilter struct {
}

func NewRawSegmentsToCarriersFilter() *SegmentsToCarriersFilter {
	return &SegmentsToCarriersFilter{}
}

func (f *SegmentsToCarriersFilter) Filter(segments <-chan RawSegment) carrier.Carriers {
	codes := make(map[string]bool)

	for s := range segments {
		codes[s.CarrierCode] = true
	}

	var ab carrier.Builder
	for code := range codes {
		ab.Append(code)
	}

	return ab.Build()
}
