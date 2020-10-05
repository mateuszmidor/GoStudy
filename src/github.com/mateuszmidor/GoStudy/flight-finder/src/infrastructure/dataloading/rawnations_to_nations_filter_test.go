package dataloading_test

import (
	"testing"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/nations"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure/dataloading"
)

func TestRawnationsFilterReturnsValidNations(t *testing.T) {
	// given
	rawNations := make(chan dataloading.RawNation, 3)
	rawNations <- dataloading.NewRawNation("CN", "CHN", "CNY", "CHINA")
	rawNations <- dataloading.NewRawNation("ES", "ESP", "EUR", "SPAIN")
	rawNations <- dataloading.NewRawNation("PL", "POL", "PLN", "POLAND")
	close(rawNations)
	// notice: expectedNations are sorted ascending
	expectedNations := nations.Nations{
		nations.NewNation("CN", "CHN", "CNY", "CHINA"),
		nations.NewNation("ES", "ESP", "EUR", "SPAIN"),
		nations.NewNation("PL", "POL", "PLN", "POLAND"),
	}

	// when
	actualNations := dataloading.FilterRawNations(rawNations)

	// then
	if len(actualNations) != len(expectedNations) {
		t.Errorf("Expected num nations %d, got %d", len(expectedNations), len(actualNations))
	}

	for i, actual := range actualNations {
		if actual != expectedNations[i] {
			t.Errorf("At index %d expected nation %v, got %v", i, expectedNations[i], actual)
		}
	}
}
