package airports_test

import (
	"airports"
	"testing"
)

func TestBuilderReturnsAllAirportsSorted(t *testing.T) {
	// given
	var b airports.Builder
	expectedAirports := airports.Airports{
		airports.NewAirport("AAA", "Andora Airport"),
		airports.NewAirport("KKK", "Kalkuta Airport"),
		airports.NewAirport("ZZZ", "Zimbabwe Airport"),
	}

	// when
	b.Append("KKK", "Kalkuta Airport")
	b.Append("AAA", "Andora Airport")
	b.Append("ZZZ", "Zimbabwe Airport")
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
