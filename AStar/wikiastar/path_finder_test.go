package wikiastar

import (
	"testing"
)

func Test1SegmentPath(t *testing.T) {
	krk := &Node{"KRK", 0, 0}
	waw := &Node{"WAW", 0, 25}

	g := NewGraph()
	g.Connect(krk, waw)

	if path := findPath(krk, waw, g); path != "KRK-WAW" {
		t.Errorf("Invalid path found: %s", path)
	}
}

func Test2SegmentPath(t *testing.T) {
	krk := &Node{"KRK", 0, 0}
	waw := &Node{"WAW", 0, 25}
	gda := &Node{"GDA", 0, 50}

	g := NewGraph()
	g.Connect(krk, waw)
	g.Connect(waw, gda)

	if path := findPath(krk, gda, g); path != "KRK-WAW-GDA" {
		t.Errorf("Invalid path found: %s", path)
	}
}

func Test2SegmentBestPath(t *testing.T) {
	krk := &Node{"KRK", 0, 0}
	waw := &Node{"WAW", 0, 25}
	wro := &Node{"WRO", -10, 25}
	gda := &Node{"GDA", 0, 50}

	g := NewGraph()
	g.Connect(krk, waw)
	g.Connect(waw, gda)
	g.Connect(krk, wro)
	g.Connect(wro, gda)

	if path := findPath(krk, gda, g); path != "KRK-WAW-GDA" {
		t.Errorf("Invalid path found: %s", path)
	}
}
