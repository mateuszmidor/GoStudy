package airports_test

import (
	"fmt"
	"testing"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/airports"
)

func TestGetByCodeReturnsValidAirport(t *testing.T) {
	// given
	// important: airports are sorted ascending for binary search
	airportList := airports.Airports{
		airports.NewAirportCodeOnly("AAA"),
		airports.NewAirportCodeOnly("KKK"),
		airports.NewAirportCodeOnly("ZZZ"),
	}
	cases := []struct {
		code string
		id   airports.ID
	}{
		{"AAA", 0},
		{"GGG", airports.NullID},
		{"KKK", 1},
		{"PPP", airports.NullID},
		{"ZZZ", 2},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Checking ID for %s", c.code), func(t *testing.T) {
			// when
			id := airportList.GetByCode(c.code)

			// then
			if id != c.id {
				t.Errorf("For %s expected ID was %d, got %d", c.code, c.id, id)
			}
		})
	}
}
