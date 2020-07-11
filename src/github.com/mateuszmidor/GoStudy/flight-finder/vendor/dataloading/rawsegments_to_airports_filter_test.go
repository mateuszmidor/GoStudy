package dataloading_test

import (
	"airport"
	"dataloading"
	"testing"
)

func TestAirportsFilterReturnsValidAirports(t *testing.T) {
	// given
	rawSegments := make(chan dataloading.RawSegment, 3)
	rawSegments <- dataloading.NewRawSegment("KRK", "WAW", "LO")
	rawSegments <- dataloading.NewRawSegment("WAW", "WRO", "LH")
	rawSegments <- dataloading.NewRawSegment("WRO", "GDN", "BY")
	close(rawSegments)
	// expected airports are sorted
	expectedAirports := airport.Airports{
		airport.NewAirport("GDN", ""),
		airport.NewAirport("KRK", ""),
		airport.NewAirport("WAW", ""),
		airport.NewAirport("WRO", ""),
	}
	filter := dataloading.NewRawSegmentsToAirportsFilter()

	// when
	actualAirports := filter.Filter(rawSegments)

	// then
	if len(expectedAirports) != len(actualAirports) {
		t.Errorf("Expected num airports %d, got %d", len(expectedAirports), len(actualAirports))
	}

	for i, actual := range actualAirports {
		if actual != expectedAirports[i] {
			t.Errorf("At index %d expected airport %v, got %v", i, expectedAirports[i], actual)
		}
	}
}
