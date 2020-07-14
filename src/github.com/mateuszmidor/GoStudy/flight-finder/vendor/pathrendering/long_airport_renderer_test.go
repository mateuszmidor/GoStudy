package pathrendering_test

import (
	"airports"
	"fmt"
	"pathrendering"
	"testing"
)

func TestRendererReturnsValidLongAirportString(t *testing.T) {
	// given
	// important: airpors are sorted ascending
	airportList := airports.Airports{
		airports.NewAirport("GDN", "Gdańsk", 0, 0),
		airports.NewAirport("KRK", "Kraków", 0, 0),
		airports.NewAirport("WAW", "Warszawa", 0, 0),
	}
	cases := []struct {
		id       airports.ID
		expected string
	}{
		{0, "Gdańsk"},
		{1, "Kraków"},
		{2, "Warszawa"},
	}

	renderer := pathrendering.NewLongAirportRenderer(airportList)

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
