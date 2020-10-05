package astext_test

import (
	"fmt"
	"testing"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/airports"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/pathrendering/astext"
)

func TestRendererReturnsValidLongAirportString(t *testing.T) {
	// given
	// important: airpors are sorted ascending
	airportList := airports.Airports{
		airports.NewAirport("GDN", "Gdańsk", "PL", 0, 0),
		airports.NewAirport("KRK", "Kraków", "PL", 0, 0),
		airports.NewAirport("WAW", "Warszawa", "PL", 0, 0),
	}
	cases := []struct {
		id       airports.ID
		expected string
	}{
		{0, "Gdańsk"},
		{1, "Kraków"},
		{2, "Warszawa"},
	}

	renderer := astext.NewLongAirportRenderer(airportList)

	for _, c := range cases {
		t.Run(fmt.Sprintf("Checking long rendered string for ID %d", c.id), func(t *testing.T) {
			// when
			actual := renderer.Render(c.id)

			// then
			if actual != c.expected {
				t.Errorf("For ID %d expected long renderer string %s, got %s", c.id, c.expected, actual)
			}
		})
	}
}
