package pathfinding

import (
	"testing"
)

func TestAddContains(t *testing.T) {
	n1 := NodeID(1)
	n2 := NodeID(2)
	n3 := NodeID(3)
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
