package nations_test

import (
	"testing"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/nations"
)

func TestBuilderReturnsAllNationsSorted(t *testing.T) {
	// given
	var b nations.Builder
	// expected nations is sorted ascending
	expectedNations := nations.Nations{
		nations.NewNation("CN", "CHN", "CNY", "CHINA"),
		nations.NewNation("ES", "ESP", "EUR", "SPAIN"),
		nations.NewNation("PL", "POL", "PLN", "POLAND"),
	}

	// when
	b.Append("ES", "ESP", "EUR", "SPAIN")
	b.Append("PL", "POL", "PLN", "POLAND")
	b.Append("CN", "CHN", "CNY", "CHINA")
	actualNations := b.Build()

	// then
	if len(expectedNations) != len(actualNations) {
		t.Fatalf("Num expected nations != num actual nations: %d : %d", len(expectedNations), len(actualNations))
	}

	for i, expected := range expectedNations {
		actual := actualNations[i]
		if actual != expected {
			t.Errorf("At index %d expected nation was %v, got %v", i, expected, actual)
		}
	}
}
