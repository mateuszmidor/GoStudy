package stringgames

import "testing"

const reverseFailureString = "Expected %q, got %q"

func TestReverseEmpty(t *testing.T) {
	in := ""
	expectedOut := ""
	actualOut := Reverse(in)
	if actualOut != expectedOut {
		t.Errorf(reverseFailureString, expectedOut, actualOut)
	}
}

func TestReverseSingleChar(t *testing.T) {
	in := "q"
	expectedOut := "q"
	actualOut := Reverse(in)
	if actualOut != expectedOut {
		t.Errorf(reverseFailureString, expectedOut, actualOut)
	}
}

func TestReverseManyChar(t *testing.T) {
	in := "abc123"
	expectedOut := "321cba"
	actualOut := Reverse(in)
	if actualOut != expectedOut {
		t.Errorf(reverseFailureString, expectedOut, actualOut)
	}
}
