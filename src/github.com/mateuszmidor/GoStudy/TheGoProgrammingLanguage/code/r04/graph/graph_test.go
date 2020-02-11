package graph

import "testing"

func TestAddEdge(t *testing.T) {
	g := NewGraph()

	if g.HasEdge("A", "B") {
		t.Error("Graph should not have edge A-B yet")
	}

	g.AddEdge("A", "B")

	if !g.HasEdge("A", "B") {
		t.Error("Graph should have edge A-B")
	}

	if g.HasEdge("B", "A") {
		t.Error("Graph should NOT have edge B-A")
	}
}
