package carrier_test

import (
	"carrier"
	"testing"
)

func TestBuilderReturnsAllCarriersSorted(t *testing.T) {
	// given
	var b carrier.Builder
	expectedCarriers := carrier.Carriers{
		carrier.NewCarrier("AA"),
		carrier.NewCarrier("BB"),
		carrier.NewCarrier("CC"),
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
