package airport_test

import (
	"airport"
	"fmt"
	"testing"
)

func TestGetByCodeReturnsValidAirport(t *testing.T) {
	// given
	// important: airports are sorted ascending for binary search
	airports := airport.Airports{
		airport.NewAirport("AAA", "Andora Airport"),
		airport.NewAirport("KKK", "Kalkuta Airport"),
		airport.NewAirport("ZZZ", "Zimbabwe Airport"),
	}
	cases := []struct {
		code string
		id   airport.AirportID
	}{
		{"AAA", 0},
		{"KKK", 1},
		{"ZZZ", 2},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Checking AirportID for %s", c.code), func(t *testing.T) {
			// when
			id := airports.GetByCode(c.code)

			// then
			if id != c.id {
				t.Errorf("For %s expected AirportID was %d, got %d", c.code, c.id, id)
			}
		})
	}
}
