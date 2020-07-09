package connections_test

import (
	"airport"
	"connections"
	"fmt"
	"segment"
	"testing"
)

func TestFindByOriginReturnsProperRange(t *testing.T) {
	// given
	segments := segment.Segments{
		segment.NewSegment(1, 10, -1),
		segment.NewSegment(1, 20, -1),
		segment.NewSegment(2, 10, -1),
		segment.NewSegment(2, 20, -1),
		segment.NewSegment(2, 30, -1),
		segment.NewSegment(3, 10, -1),
		segment.NewSegment(3, 20, -1),
	}

	cases := []struct {
		id          airport.AirportID
		first, last segment.ID
	}{
		{1, 0, 2},
		{2, 2, 5},
		{3, 5, 7},
		{4, 7, 7},
	}

	var finder connections.SegmentRangeFinder
	for _, c := range cases {
		t.Run(fmt.Sprintf("Searching airportID = %d", c.id), func(t *testing.T) {
			// when
			first, last := finder.ByFromAirport(segments, c.id)

			// then
			if first != c.first {
				t.Errorf("Expected first = %d, got %d", c.first, first)
			}
			if last != c.last {
				t.Errorf("Expected last = %d, got %d", c.last, last)
			}
		})
	}
}
