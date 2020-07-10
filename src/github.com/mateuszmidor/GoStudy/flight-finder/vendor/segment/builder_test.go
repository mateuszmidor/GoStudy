package segment_test

import (
	"airport"
	"segment"
	"testing"
)

func TestBuilderReturnsAllSegmentsSorted(t *testing.T) {
	// given
	airports := airport.Airports{
		airport.NewAirport("AAA", "Andora Airport"),   // ID 0
		airport.NewAirport("KKK", "Kalkuta Airport"),  // ID 1
		airport.NewAirport("ZZZ", "Zimbabwe Airport"), // ID 2
	}
	carriers := carrier.Carriers{
		carrier.NewCarrier("LH"),
		carrier.NewCarrier("LO"),
	}
	expectedSegments := []segment.Segment{
		segment.NewSegment(0, 1, 1),
		segment.NewSegment(1, 2, 0),
	}

	// when
	b := segment.NewBuilder(airports)
	b.Append("KKK", "ZZZ", "LH")
	b.Append("AAA", "KKK", "LO")
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
