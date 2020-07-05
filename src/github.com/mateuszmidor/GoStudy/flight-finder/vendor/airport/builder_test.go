package airport_test

import (
	"airport"
	"testing"
)

func TestBuilderReturnsAllAirportsSorted(t *testing.T) {
	// given
	var b airport.Builder
	expectedAirports := airport.Airports{
		airport.NewAirport("AAA", "Andora Airport"),
		airport.NewAirport("KKK", "Kalkuta Airport"),
		airport.NewAirport("ZZZ", "Zimbabwe Airport"),
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
