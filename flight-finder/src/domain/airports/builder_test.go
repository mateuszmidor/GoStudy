package airports_test

import (
	"testing"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/airports"
)

func TestBuilderReturnsAllAirportsSorted(t *testing.T) {
	// given
	var b airports.Builder
	expectedAirports := airports.Airports{
		airports.NewAirportCodeOnly("AAA"),
		airports.NewAirportCodeOnly("KKK"),
		airports.NewAirportCodeOnly("ZZZ"),
	}

	// when
	b.Append("KKK")
	b.Append("AAA")
	b.Append("ZZZ")
	actualAirports := b.Build()

	// then
	if len(expectedAirports) != len(actualAirports) {
		t.Fatalf("Num expected airports != num actual airports: %d : %d", len(expectedAirports), len(actualAirports))
	}

	for i, expected := range expectedAirports {
		actual := actualAirports[i]
		if actual != expected {
			t.Errorf("At index %d expected airport %v, got %v", i, expected, actual)
		}
	}
}
