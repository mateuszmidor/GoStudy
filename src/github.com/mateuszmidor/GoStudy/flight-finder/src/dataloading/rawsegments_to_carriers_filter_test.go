package dataloading_test

import (
	"testing"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/carriers"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/dataloading"
)

func TestAirportsCarriersLoaderReturnsValidData2(t *testing.T) {
	// given
	rawSegments := make(chan dataloading.RawSegment, 3)
	rawSegments <- dataloading.NewRawSegment("KRK", "WAW", "LO")
	rawSegments <- dataloading.NewRawSegment("WAW", "WRO", "LH")
	rawSegments <- dataloading.NewRawSegment("WRO", "GDN", "BY")
	close(rawSegments)
	// expected carriers are sorted
	expectedCarrires := carriers.Carriers{
		carriers.NewCarrier("BY"),
		carriers.NewCarrier("LH"),
		carriers.NewCarrier("LO"),
	}

	// when
	actualCarriers := dataloading.FilterCarriers(rawSegments)

	// then
	if len(expectedCarrires) != len(actualCarriers) {
		t.Errorf("Expected num carriers %d, got %d", len(expectedCarrires), len(actualCarriers))
	}

	for i, actual := range actualCarriers {
		if actual != actualCarriers[i] {
			t.Errorf("At index %d expected carrier %v, got %v", i, expectedCarrires[i], actual)
		}
	}
}
