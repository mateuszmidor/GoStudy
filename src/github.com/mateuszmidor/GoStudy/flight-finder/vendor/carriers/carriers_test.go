package carriers_test

import (
	"carriers"
	"fmt"
	"testing"
)

func TestGetByCodeReturnsValidCarrier(t *testing.T) {
	// given
	// important: carriers are sorted ascending for binary search
	carrierList := carriers.Carriers{
		carriers.NewCarrier("AA"),
		carriers.NewCarrier("KK"),
		carriers.NewCarrier("ZZ"),
	}
	cases := []struct {
		code string
		id   carriers.ID
	}{
		{"AA", 0},
		{"GG", carriers.NullID},
		{"KK", 1},
		{"PP", carriers.NullID},
		{"ZZ", 2},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Checking CarrierID for %s", c.code), func(t *testing.T) {
			// when
			id := carrierList.GetByCode(c.code)

			// then
			if id != c.id {
				t.Errorf("For %s expected CarrierID was %d, got %d", c.code, c.id, id)
			}
		})
	}
}
