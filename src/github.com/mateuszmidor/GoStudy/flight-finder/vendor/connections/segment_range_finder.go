package connections

import (
	"airport"
	"segment"
	"sort"
)

type SegmentRangeFinder struct {
}

func (s *SegmentRangeFinder) ByFromAirport(segments segment.Segments, id airport.ID) (first, last segment.ID) {
	lo, hi := lowerBound(segments, id), upperBound(segments, id)
	if lo == hi {
		return segment.NullID, segment.NullID
	}
	return segment.ID(lo), segment.ID(hi)
}

// lowerBound is index of first matching element
func lowerBound(segments segment.Segments, id airport.ID) int {
	ge := func(i int) bool {
		return segments[i].From() >= id
	}

	return sort.Search(len(segments), ge)
}

// upperBound is index of last matching element +1
func upperBound(segments segment.Segments, id airport.ID) int {
	g := func(i int) bool {
		return segments[i].From() > id
	}

	return sort.Search(len(segments), g)
}
