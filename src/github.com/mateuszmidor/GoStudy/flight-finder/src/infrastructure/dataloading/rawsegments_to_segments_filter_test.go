package dataloading_test

import (
	"testing"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/airports"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/carriers"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/segments"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure/dataloading"
)

func TestRawsegmentsFilterReturnsValidSegments(t *testing.T) {
	// given
	// important: airports are sorted
	airports := airports.Airports{
		airports.NewAirportCodeOnly("GDN"), // ID=0
		airports.NewAirportCodeOnly("KRK"), // ID=1
		airports.NewAirportCodeOnly("WAW"), // ID=2
		airports.NewAirportCodeOnly("WRO"), // ID=3
	}
	// important: carrierrs are sorted
	carriers := carriers.Carriers{
		carriers.NewCarrier("BY"), // carrierID=0
		carriers.NewCarrier("LH"), // carrierID=1
		carriers.NewCarrier("LO"), // carrierID=2
	}
	// important: expected segments are sorted
	expectedSegments := segments.Segments{
		segments.NewSegment(1, 2, 2),
		segments.NewSegment(2, 3, 1),
		segments.NewSegment(3, 0, 0),
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
