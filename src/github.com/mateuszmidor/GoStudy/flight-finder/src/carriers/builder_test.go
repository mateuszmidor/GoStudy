package carriers_test

import (
	"testing"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/carriers"
)

func TestBuilderReturnsAllCarriersSorted(t *testing.T) {
	// given
	var b carriers.Builder
	expectedCarriers := carriers.Carriers{
		carriers.NewCarrier("AA"),
		carriers.NewCarrier("BB"),
		carriers.NewCarrier("CC"),
	}

	// when
	b.Append("AA")
	b.Append("BB")
	b.Append("CC")
	actualCarriers := b.Build()

	// then
	if len(expectedCarriers) != len(actualCarriers) {
		t.Fatalf("Num expected airports != num actual airports: %d : %d", len(expectedCarriers), len(actualCarriers))
	}

	for i, expected := range expectedCarriers {
		actual := actualCarriers[i]
		if actual != expected {
			t.Errorf("At index %d expected carrier %v, got %v", i, expected, actual)
		}
	}
}
