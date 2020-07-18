package astext_test

import (
	"airports"
	"fmt"
	"pathrendering/astext"
	"testing"
)

func TestRendererReturnsValidShortAirportString(t *testing.T) {
	// given
	// important: airpors are sorted ascending
	airportList := airports.Airports{
		airports.NewAirport("GDN", "", 0, 0),
		airports.NewAirport("KRK", "", 0, 0),
		airports.NewAirport("WAW", "", 0, 0),
	}
	cases := []struct {
		id       airports.ID
		expected string
	}{
		{0, "GDN"},
		{1, "KRK"},
		{2, "WAW"},
	}

	renderer := astext.NewShortAirportRenderer(airportList)

	for _, c := range cases {
		t.Run(fmt.Sprintf("Checking short rendered string for ID %d", c.id), func(t *testing.T) {
			// when
			actual := renderer.Render(c.id)

			// then
			if actual != c.expected {
				t.Errorf("For ID %d expected short renderer string %s, got %s", c.id, c.expected, actual)
			}
		})
	}
}
