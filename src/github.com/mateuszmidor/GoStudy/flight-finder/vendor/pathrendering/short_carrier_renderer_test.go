package pathrendering_test

import (
	"carrier"
	"fmt"
	"pathrendering"
	"testing"
)

func TestRendererReturnsValidShortCarrierString(t *testing.T) {
	// given
	// important: carriers are sorted ascending
	carriers := carrier.Carriers{
		carrier.NewCarrier("AA"), // id=0
		carrier.NewCarrier("BB"), // id=1
		carrier.NewCarrier("CC"), // id=2
	}
	cases := []struct {
		id       carrier.ID
		expected string
	}{
		{0, "AA"},
		{1, "BB"},
		{2, "CC"},
	}
	renderer := pathrendering.NewShortCarrierRenderer(carriers)

	for _, c := range cases {
		t.Run(fmt.Sprintf("Checking short rendered string for CarrierID %d", c.id), func(t *testing.T) {
			// when
			actual := renderer.Render(c.id)

			// then
			if actual != c.expected {
				t.Errorf("For CarrierID %d expected short renderer string %s, got %s", c.id, c.expected, actual)
			}
		})
	}
}
