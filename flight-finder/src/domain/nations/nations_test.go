package nations_test

import (
	"fmt"
	"testing"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/nations"
)

func TestGetByCodeReturnsValidNation(t *testing.T) {
	// given
	// important: nation are sorted ascending for binary search
	nationList := nations.Nations{
		nations.NewNation("CN", "CHN", "CNY", "CHINA"),
		nations.NewNation("ES", "ESP", "EUR", "SPAIN"),
		nations.NewNation("PL", "POL", "PLN", "POLAND"),
	}
	cases := []struct {
		code string
		id   nations.ID
	}{
		{"CN", 0},
		{"GG", nations.NullID},
		{"ES", 1},
		{"GG", nations.NullID},
		{"PL", 2},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Checking ID for %s", c.code), func(t *testing.T) {
			// when
			id := nationList.GetByCode(c.code)

			// then
			if id != c.id {
				t.Errorf("For %s expected ID was %d, got %d", c.code, c.id, id)
			}
		})
	}
}
