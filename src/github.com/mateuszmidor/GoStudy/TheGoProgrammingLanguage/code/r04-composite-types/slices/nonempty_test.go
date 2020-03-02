package slices

import (
	"strings"
	"testing"
)

type data struct {
	in  []string
	out []string
}

const nonEmptyFailureString = "%v != %v"

func TestNonEmpty(t *testing.T) {
	in := []data{
		{[]string{"one", "two", "three"}, []string{"one", "two", "three"}},
		{[]string{"", "two", "three"}, []string{"two", "three"}},
		{[]string{"one", "", "three"}, []string{"one", "three"}},
		{[]string{"one", "two", ""}, []string{"one", "two"}},
		{[]string{"", "two", ""}, []string{"two"}},
		{[]string{"", "", ""}, []string{}},
	}

	for _, tc := range in {
		if actualOut := nonEmpty(tc.in); slicesDifferent(actualOut, tc.out) {
			t.Errorf(nonEmptyFailureString, actualOut, tc.out)
		}
	}
}

func slicesDifferent(a []string, b []string) bool {
	return strings.Join(a, " ") != strings.Join(b, " ")
}
