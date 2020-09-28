package dataloading_test

import (
	"testing"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/airports"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/dataloading"
)

func TestAirportsFilterReturnsValidAirports(t *testing.T) {
	// given
	rawSegments := make(chan dataloading.RawSegment, 3)
	rawSegments <- dataloading.NewRawSegment("KRK", "WAW", "LO")
	rawSegments <- dataloading.NewRawSegment("WAW", "WRO", "LH")
	rawSegments <- dataloading.NewRawSegment("WRO", "GDN", "BY")
	close(rawSegments)
	// expected airports are sorted
	expectedAirports := airports.Airports{
		airports.NewAirportCodeOnly("GDN"),
		airports.NewAirportCodeOnly("KRK"),
		airports.NewAirportCodeOnly("WAW"),
		airports.NewAirportCodeOnly("WRO"),
	}

	// when
	actualAirports := dataloading.FilterAirports(rawSegments)

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
