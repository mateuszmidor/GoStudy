package segment_test

import (
	"fmt"
	"segment"
	"strings"
	"testing"
)

func TestLoadValidCSVShouldReturnAllSegments(t *testing.T) {
	// given
	var source segment.SourceCSV
	actualSegments := make(chan segment.RawSegment, 1)
	expectedSegments := []segment.RawSegment{
		{"GDY", "WAW", "BY"},
		{"WAW", "KRK", "LH"},
		{"KRK", "KTW", "LO"},
	}
	csv := `
"GDY","WAW","BY"
"WAW","KRK","LH"
"KRK","KTW","LO"
`
	// when
	go source.StartLoadingSegments(strings.NewReader(csv), actualSegments)

	// then
	errorDetails := checkExpectedSegments(expectedSegments, actualSegments)
	if errorDetails != "" {
		t.Error(errorDetails)
	}
}

func TestLoadBrokenCSVShouldReturnOnlyValidSegments(t *testing.T) {
	// given
	var source segment.SourceCSV
	actualSegments := make(chan segment.RawSegment, 1)
	expectedSegments := []segment.RawSegment{
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
	go source.StartLoadingSegments(strings.NewReader(csv), actualSegments)

	// then
	errorDetails := checkExpectedSegments(expectedSegments, actualSegments)
	if errorDetails != "" {
		t.Error(errorDetails)
	}
}

func checkExpectedSegments(expected []segment.RawSegment, actual chan segment.RawSegment) string {
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

func findSegment(subject segment.RawSegment, list []segment.RawSegment) int {
	for i, seg := range list {
		if seg == subject {
			return i
		}
	}
	return -1
}

func removeSegment(index int, list *[]segment.RawSegment) {
	l := *list
	l[index] = l[len(l)-1]
	l = l[:len(l)-1]
	*list = l
}
