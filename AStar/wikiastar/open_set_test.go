package wikiastar

import (
	"testing"
)

func TestAddRemoveIsEmpty(t *testing.T) {
	n1 := &Node{}
	n2 := &Node{}
	n3 := &Node{}
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
