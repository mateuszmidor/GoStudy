package pathrendering_test

import (
	"carriers"
	"fmt"
	"pathrendering"
	"testing"
)

func TestRendererReturnsValidShortCarrierString(t *testing.T) {
	// given
	// important: carriers are sorted ascending
	carrierList := carriers.Carriers{
		carriers.NewCarrier("AA"), // id=0
		carriers.NewCarrier("BB"), // id=1
		carriers.NewCarrier("CC"), // id=2
	}
	cases := []struct {
		id       carriers.ID
		expected string
	}{
		{0, "AA"},
		{1, "BB"},
		{2, "CC"},
	}
	renderer := pathrendering.NewShortCarrierRenderer(carrierList)

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
