package multipathastar

import (
	"testing"
)

func TestGetSet(t *testing.T) {
	n1 := NodeID(1)
	n2 := NodeID(2)
	n3 := NodeID(3)
	score := NewScore()

	score.Set(n1, 10)
	score.Set(n2, 20)
	score.Set(n3, 30)

	if score.Get(n1) != 10 {
		t.Error("Score for n1 should be 10")
	}

	if score.Get(n2) != 20 {
		t.Error("Score for n1 should be 20")
	}

	if score.Get(n3) != 30 {
		t.Error("Score for n1 should be 30")
	}
}

func TestDefaultValueIsInfinity(t *testing.T) {
	n1 := NodeID(1)
	n2 := NodeID(2)
	n3 := NodeID(3)
	n4 := NodeID(4)
	score := NewScore()

	score.Set(n1, 10)
	score.Set(n2, 20)
	score.Set(n3, 30)
	// n4 not set

	if score.Get(n4) != score.GetInfinity() {
		t.Error("Default score should be +infinity")
	}

}
