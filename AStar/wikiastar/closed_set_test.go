package wikiastar

import (
	"testing"
)

func TestAddContains(t *testing.T) {
	n1 := &Node{}
	n2 := &Node{}
	n3 := &Node{}
	cs := NewClosedSet()

	if cs.Contains(n1) || cs.Contains(n2) || cs.Contains(n3) {
		t.Error("ClosedSet should not contain any elements")
	}

	cs.Add(n1)
	if !cs.Contains(n1) {
		t.Error("ClosedSet should contain element n1")
	}

	cs.Add(n2)
	if !cs.Contains(n2) {
		t.Error("ClosedSet should contain element n2")
	}

	cs.Add(n3)
	if !cs.Contains(n3) {
		t.Error("ClosedSet should contain element n3")
	}

	if !(cs.Contains(n1) && cs.Contains(n2) && cs.Contains(n3)) {
		t.Error("ClosedSet should contain elements n1, n2, n3")
	}
}
