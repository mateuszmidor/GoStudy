package slices

import "testing"

type twostrings struct {
	in  string
	out string
}

const squashSpacesFailureString = "(actual) %q != %q (expected)"

func TestSquashSpaces(t *testing.T) {
	in := []twostrings{
		{"", ""},
		{"a", "a"},
		{"ab", "ab"},
		{"a b", "a b"},
		{"a  b", "a b"},
		{"a   b", "a b"},
		{" a   b ", " a b "},
	}

	for _, tc := range in {
		if actualOut := squashSpaces(tc.in); actualOut != tc.out {
			t.Errorf(squashSpacesFailureString, actualOut, tc.out)
		}
	}
}
