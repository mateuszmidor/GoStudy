package wikiastar

import (
	"testing"
)

func TestConnectEstablishedUnidirectionalNeighbors(t *testing.T) {
	krk := &Node{}
	waw := &Node{}
	wro := &Node{}
	gda := &Node{}

	g := NewGraph()
	g.Connect(krk, waw)
	g.Connect(waw, gda)
	g.Connect(krk, wro)
	g.Connect(wro, gda)

	if _, ok := g.GetNeighbors(krk)[waw]; !ok {
		t.Error("Missing edge krk-waw")
	}

	if _, ok := g.GetNeighbors(waw)[gda]; !ok {
		t.Error("Missing edge waw-gda")
	}

	if _, ok := g.GetNeighbors(krk)[wro]; !ok {
		t.Error("Missing edge krk-waw")
	}

	if _, ok := g.GetNeighbors(wro)[gda]; !ok {
		t.Error("Missing edge waw-gda")
	}

	// this edge should not exist
	if _, ok := g.GetNeighbors(krk)[gda]; ok {
		t.Error("Unexpected edge krk-gda")
	}

	// this edge should not exist
	if _, ok := g.GetNeighbors(gda)[wro]; ok {
		t.Error("Unexpected edge gda-wro")
	}

	// this edge should not exist
	if _, ok := g.GetNeighbors(wro)[krk]; ok {
		t.Error("Unexpected edge wro-krk")
	}
	// this edge should not exist
	if _, ok := g.GetNeighbors(gda)[wro]; ok {
		t.Error("Unexpected edge gda-wro")
	}

	// this edge should not exist
	if _, ok := g.GetNeighbors(wro)[krk]; ok {
		t.Error("Unexpected edge wro-krk")
	}

	// this edge should not exist
	if _, ok := g.GetNeighbors(gda)[waw]; ok {
		t.Error("Unexpected edge gda-waw")
	}

	// this edge should not exist
	if _, ok := g.GetNeighbors(waw)[krk]; ok {
		t.Error("Unexpected edge waw-krk")
	}

	// this edge should not exist
	if _, ok := g.GetNeighbors(waw)[wro]; ok {
		t.Error("Unexpected edge waw-wro")
	}

	// this edge should not exist
	if _, ok := g.GetNeighbors(wro)[waw]; ok {
		t.Error("Unexpected edge wro-waw")
	}
}
