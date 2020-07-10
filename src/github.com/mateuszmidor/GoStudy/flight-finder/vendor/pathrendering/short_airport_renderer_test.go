package pathrendering_test

import (
	"airport"
	"fmt"
	"pathrendering"
	"testing"
)

func TestRendererReturnsValidShortAirportString(t *testing.T) {
	// given
	// important: airpors are sorted ascending
	airports := airport.Airports{
		airport.NewAirport("GDN", ""),
		airport.NewAirport("KRK", ""),
		airport.NewAirport("WAW", ""),
	}
	cases := []struct {
		id       airport.ID
		expected string
	}{
		{0, "GDN"},
		{1, "KRK"},
		{2, "WAW"},
	}

	renderer := pathrendering.NewShortAirportRenderer(airports)

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
