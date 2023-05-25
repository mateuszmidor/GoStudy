package wikiastar

import (
	"testing"
)

func TestGetSet(t *testing.T) {
	n1 := &Node{}
	n2 := &Node{}
	n3 := &Node{}
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
	n1 := &Node{}
	n2 := &Node{}
	n3 := &Node{}
	n4 := &Node{}
	score := NewScore()

	score.Set(n1, 10)
	score.Set(n2, 20)
	score.Set(n3, 30)
	// n4 not set

	if score.Get(n4) != score.GetInfinity() {
		t.Error("Default score should be +infinity")
	}

}
