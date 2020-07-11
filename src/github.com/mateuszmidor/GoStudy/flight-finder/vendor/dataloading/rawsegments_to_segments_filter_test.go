package dataloading_test

import (
	"airport"
	"carrier"
	"dataloading"
	"segment"
	"testing"
)

func TestRawsegmentsFilterReturnsValidSegments(t *testing.T) {
	// given
	// important: airports are sorted
	airports := airport.Airports{
		airport.NewAirport("GDN", ""), // ID=0
		airport.NewAirport("KRK", ""), // ID=1
		airport.NewAirport("WAW", ""), // ID=2
		airport.NewAirport("WRO", ""), // ID=3
	}
	// important: carrierrs are sorted
	carriers := carrier.Carriers{
		carrier.NewCarrier("BY"), // carrierID=0
		carrier.NewCarrier("LH"), // carrierID=1
		carrier.NewCarrier("LO"), // carrierID=2
	}
	// important: expected segments are sorted
	expectedSegments := segment.Segments{
		segment.NewSegment(1, 2, 2),
		segment.NewSegment(2, 3, 1),
		segment.NewSegment(3, 0, 0),
	}
	rawSegments := make(chan dataloading.RawSegment, 3)
	rawSegments <- dataloading.NewRawSegment("KRK", "WAW", "LO")
	rawSegments <- dataloading.NewRawSegment("WAW", "WRO", "LH")
	rawSegments <- dataloading.NewRawSegment("WRO", "GDN", "BY")
	close(rawSegments)
	filter := dataloading.NewRawSegmentsToSegmentsFilter(airports, carriers)

	// when
	actualSegments := filter.Filter(rawSegments)

	// then
	if len(actualSegments) != len(expectedSegments) {
		t.Errorf("Expected num segments %d, got %d", len(expectedSegments), len(actualSegments))
	}

	for i, actual := range actualSegments {
		if actual != expectedSegments[i] {
			t.Errorf("At index %d expected segment %v, got %v", i, expectedSegments[i], actual)
		}
	}
}
