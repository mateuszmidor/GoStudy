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
		segment.NewSegment(4, 10, -1),
		segment.NewSegment(4, 20, -1),
	}

	cases := []struct {
		id          airport.ID
		first, last segment.ID
	}{
		{1, 0, 2},
		{2, 2, 5},
		{3, segment.NullID, segment.NullID},
		{4, 5, 7},
		{5, segment.NullID, segment.NullID},
	}

	var finder connections.SegmentRangeFinder
	for _, c := range cases {
		t.Run(fmt.Sprintf("Searching ID = %d", c.id), func(t *testing.T) {
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
