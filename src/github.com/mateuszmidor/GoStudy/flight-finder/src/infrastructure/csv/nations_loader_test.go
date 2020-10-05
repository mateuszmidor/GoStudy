package csv_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure/csv"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure/dataloading"
)

func TestLoadValidCSVShouldReturnAllNations(t *testing.T) {
	// given
	var loader csv.NationsLoader
	actualNations := make(chan dataloading.RawNation, 1)
	expectedNations := []dataloading.RawNation{
		dataloading.NewRawNation("CN", "CHN", "CNY", "CHINA"),
		dataloading.NewRawNation("ES", "ESP", "EUR", "SPAIN"),
		dataloading.NewRawNation("PL", "POL", "PLN", "POLAND"),
	}
	// NATION,ISO,CURRENCY,DESCRIPTION
	csv := `
"CN","CHN","CNY","CHINA"
"ES","ESP","EUR","SPAIN"
"PL","POL","PLN","POLAND"
`
	// when
	go loader.StartLoading(strings.NewReader(csv), actualNations)

	// then
	errorDetails := checkExpectedNations(expectedNations, actualNations)
	if errorDetails != "" {
		t.Error(errorDetails)
	}
}

func checkExpectedNations(expected []dataloading.RawNation, actual chan dataloading.RawNation) string {
	var result string
	for seg := range actual {
		if index := findNation(seg, expected); index != -1 {
			removeNation(index, &expected)
		} else {
			result += fmt.Sprintf("Unexpected Nation loaded: %+v\n", seg)
		}
	}

	if len(expected) != 0 {
		result += fmt.Sprintf("Expected Nations not loaded: %+v", expected)
	}

	return result
}

func findNation(subject dataloading.RawNation, list []dataloading.RawNation) int {
	for i, seg := range list {
		if seg == subject {
			return i
		}
	}
	return -1
}

func removeNation(index int, list *[]dataloading.RawNation) {
	l := *list
	l[index] = l[len(l)-1]
	l = l[:len(l)-1]
	*list = l
}
