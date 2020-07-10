package carrier_test

import (
	"carrier"
	"fmt"
	"testing"
)

func TestGetByCodeReturnsValidCarrier(t *testing.T) {
	// given
	// important: carriers are sorted ascending for binary search
	carriers := carrier.Carriers{
		carrier.NewCarrier("AA"),
		carrier.NewCarrier("KK"),
		carrier.NewCarrier("ZZ"),
	}
	cases := []struct {
		code string
		id   carrier.ID
	}{
		{"AA", 0},
		{"GG", carrier.NullID},
		{"KK", 1},
		{"PP", carrier.NullID},
		{"ZZ", 2},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Checking CarrierID for %s", c.code), func(t *testing.T) {
			// when
			id := carriers.GetByCode(c.code)

			// then
			if id != c.id {
				t.Errorf("For %s expected CarrierID was %d, got %d", c.code, c.id, id)
			}
		})
	}
}
