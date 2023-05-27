package multipathastar

import (
	"testing"
)

func TestAddRemoveIsEmpty(t *testing.T) {
	n1 := NodeID(1)
	n2 := NodeID(2)
	n3 := NodeID(3)
	os := NewOpenSet()

	if !os.IsEmpty() {
		t.Error("OpenSet should be empty")
	}

	os.Add(n1)
	os.Add(n2)
	if os.IsEmpty() {
		t.Error("OpenSet should NOT be empty")
	}

	os.Remove(n3)
	os.Remove(n2)
	os.Remove(n1)
	if !os.IsEmpty() {
		t.Error("OpenSet should be empty")
	}
}
