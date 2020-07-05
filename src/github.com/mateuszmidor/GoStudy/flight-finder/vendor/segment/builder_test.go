package segment_test

import (
	"airport"
	"segment"
	"testing"
)

func TestBuilderReturnsAllSegmentsSorted(t *testing.T) {
	// given
	airports := airport.Airports{
		airport.NewAirport("AAA", "Andora Airport"),   // AirportID 0
		airport.NewAirport("KKK", "Kalkuta Airport"),  // AirportID 1
		airport.NewAirport("ZZZ", "Zimbabwe Airport"), // AirportID 2
	}
	b := segment.NewBuilder(airports)
	expectedSegments := []segment.Segment{
		segment.NewSegment(0, 1),
		segment.NewSegment(1, 2),
	}

	// when
	b.Append("KKK", "ZZZ")
	b.Append("AAA", "KKK")
	actualSegments := b.Build()

	// then
	if len(expectedSegments) != len(actualSegments) {
		t.Fatalf("Num expected segments != num actual segments: %d : %d", len(expectedSegments), len(actualSegments))
	}

	for i, expected := range expectedSegments {
		actual := actualSegments[i]
		if actual != expected {
			t.Errorf("At index %d expected segment %v, got %v", i, expected, actual)
		}
	}
}
