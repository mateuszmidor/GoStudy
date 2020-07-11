package dataloading

import (
	"carriers"
)

type SegmentsToCarriersFilter struct {
}

func NewRawSegmentsToCarriersFilter() *SegmentsToCarriersFilter {
	return &SegmentsToCarriersFilter{}
}

func (f *SegmentsToCarriersFilter) Filter(segments <-chan RawSegment) carriers.Carriers {
	codes := make(map[string]bool)

	for s := range segments {
		codes[s.CarrierCode] = true
	}

	var ab carriers.Builder
	for code := range codes {
		ab.Append(code)
	}

	return ab.Build()
}
