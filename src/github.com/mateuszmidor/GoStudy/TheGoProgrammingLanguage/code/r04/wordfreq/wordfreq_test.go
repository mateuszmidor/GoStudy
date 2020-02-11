package main

import (
	"fmt"
	"strings"
	"testing"
)

func checkFreq(actual, expected map[string]int) (result string) {
	for word, actualcount := range actual {
		expectedcount := expected[word]
		if actualcount != expectedcount {
			result += fmt.Sprintf("for word %q actual count=%d, expected count=%d\n", word, actualcount, expectedcount)
		}
	}

	for word, expectedcount := range expected {
		actualcount := actual[word]
		if actualcount != expectedcount {
			result += fmt.Sprintf("for word %q actual count=%d, expected count=%d\n", word, actualcount, expectedcount)
		}
	}

	return
}
func TestEmptyInput(t *testing.T) {
	s := ""
	input := strings.NewReader(s)
	freq := Wordfreq(input)

	result := checkFreq(freq, map[string]int{})
	if result != "" {
		t.Errorf("Result: \n%s\n", result)
	}
}

func TestSingleWord(t *testing.T) {
	s := "carburator"
	input := strings.NewReader(s)
	freq := Wordfreq(input)

	result := checkFreq(freq, map[string]int{"carburator": 1})
	if result != "" {
		t.Errorf("Result: \n%s\n", result)
	}
}

func TestDoubleWord(t *testing.T) {
	s := "carburator carburator"
	input := strings.NewReader(s)
	freq := Wordfreq(input)

	result := checkFreq(freq, map[string]int{"carburator": 2})
	if result != "" {
		t.Errorf("Result: \n%s\n", result)
	}
}

func TestMultipleWords(t *testing.T) {
	s := "carburator break carburator headlight"
	input := strings.NewReader(s)
	freq := Wordfreq(input)

	result := checkFreq(freq, map[string]int{"carburator": 2, "break": 1, "headlight": 1})
	if result != "" {
		t.Errorf("Result: \n%s\n", result)
	}
}
