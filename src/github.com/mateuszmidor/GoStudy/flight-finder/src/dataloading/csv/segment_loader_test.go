package csv_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/dataloading"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/dataloading/csv"
)

func TestLoadValidCSVShouldReturnAllSegments(t *testing.T) {
	// given
	var loader csv.SegmentLoader
	actualSegments := make(chan dataloading.RawSegment, 1)
	expectedSegments := []dataloading.RawSegment{
		{FromAirportCode: "GDY", ToAirportCode: "WAW", CarrierCode: "BY"},
		{FromAirportCode: "WAW", ToAirportCode: "KRK", CarrierCode: "LH"},
		{FromAirportCode: "KRK", ToAirportCode: "KTW", CarrierCode: "LO"},
	}
	csv := `
"GDY","WAW","BY"
"WAW","KRK","LH"
"KRK","KTW","LO"
`
	// when
	go loader.StartLoading(strings.NewReader(csv), actualSegments)

	// then
	errorDetails := checkExpectedSegments(expectedSegments, actualSegments)
	if errorDetails != "" {
		t.Error(errorDetails)
	}
}

func TestLoadBrokenCSVShouldReturnOnlyValidSegments(t *testing.T) {
	// given
	var loader csv.SegmentLoader
	actualSegments := make(chan dataloading.RawSegment, 1)
	expectedSegments := []dataloading.RawSegment{
		// {"GDY", "WAW", "BY"},
		// {"WAW", "KRK", "LH"},
		{"KRK", "KTW", "LO"},
	}
	csv := `
"GDY","WAW", "BY"
"WAW","KRK"
"KRK","KTW","LO"
`
	// when
	go loader.StartLoading(strings.NewReader(csv), actualSegments)

	// then
	errorDetails := checkExpectedSegments(expectedSegments, actualSegments)
	if errorDetails != "" {
		t.Error(errorDetails)
	}
}

func checkExpectedSegments(expected []dataloading.RawSegment, actual chan dataloading.RawSegment) string {
	var result string
	for seg := range actual {
		if index := findSegment(seg, expected); index != -1 {
			removeSegment(index, &expected)
		} else {
			result += fmt.Sprintf("Unexpected segment loaded: %v\n", seg)
		}
	}

	if len(expected) != 0 {
		result += fmt.Sprintf("Expected segments not loaded: %v", expected)
	}

	return result
}

func findSegment(subject dataloading.RawSegment, list []dataloading.RawSegment) int {
	for i, seg := range list {
		if seg == subject {
			return i
		}
	}
	return -1
}

func removeSegment(index int, list *[]dataloading.RawSegment) {
	l := *list
	l[index] = l[len(l)-1]
	l = l[:len(l)-1]
	*list = l
}
