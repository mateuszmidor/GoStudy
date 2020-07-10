package dataloading_test

import (
	"carrier"
	"dataloading"
	"segment"
	"testing"
)

func TestAirportsCarriersLoaderReturnsValidData2(t *testing.T) {
	// given
	rawSegments := make(chan segment.RawSegment, 3)
	rawSegments <- segment.NewRawSegment("KRK", "WAW", "LO")
	rawSegments <- segment.NewRawSegment("WAW", "WRO", "LH")
	rawSegments <- segment.NewRawSegment("WRO", "GDN", "BY")
	close(rawSegments)
	// expected carriers are sorted
	expectedCarrires := carrier.Carriers{
		carrier.NewCarrier("BY"),
		carrier.NewCarrier("LH"),
		carrier.NewCarrier("LO"),
	}
	filter := dataloading.NewRawSegmentsToCarriersFilter()

	// when
	actualCarriers := filter.Filter(rawSegments)

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
