package connections

import (
	"airport"
	"segment"
)

type SegmentRangeFinder struct {
}

func (s *SegmentRangeFinder) ByFromAirport(segments segment.Segments, id airport.AirportID) (first, last segment.ID) {
	return segment.ID(lowerBound(segments, id)), segment.ID(upperBound(segments, id))
}

// lowerBound is index of first matching element
func lowerBound(segments segment.Segments, id airport.AirportID) int {
	first := 0
	last := len(segments)
	count := last - first

	for count > 0 {
		i := first
		step := count / 2
		i += step
		if segments[i].From() < id {
			first = i + 1
			count -= step + 1
		} else {
			count = step
		}
	}
	return first
}

// upperBound is index of last matching element +1
func upperBound(segments segment.Segments, id airport.AirportID) int {
	first := 0
	last := len(segments)
	count := last - first

	for count > 0 {
		i := first
		step := count / 2
		i += step
		if segments[i].From() <= id {
			first = i + 1
			count -= step + 1
		} else {
			count = step
		}
	}
	return first
}
