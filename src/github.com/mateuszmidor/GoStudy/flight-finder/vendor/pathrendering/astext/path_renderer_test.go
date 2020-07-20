package astext_test

import (
	"airports"
	"bytes"
	"carriers"
	"pathfinding"
	"pathrendering/astext"
	"segments"
	"strconv"
	"testing"
)

type stubAirportRenderer struct {
}

func (r *stubAirportRenderer) Render(id airports.ID) string {
	return strconv.Itoa(int(id))
}

type stubCarrierRenderer struct {
}

func (r *stubCarrierRenderer) Render(id carriers.ID) string {
	return strconv.Itoa(int(id))
}

func TestPathRendererTurnsEmptyPathsIntoEmptyPathString(t *testing.T) {
	// given
	segments := segments.Segments{}
	emptyPaths := []pathfinding.Path{}
	expected := ""
	airportRenderer := stubAirportRenderer{}
	carrierRenderer := stubCarrierRenderer{}
	pathRenderer := astext.NewPathRenderer(&airportRenderer, &carrierRenderer, segments, ",")
	buf := bytes.NewBuffer([]byte{})

	// when
	pathRenderer.Render(buf, emptyPaths)
	actual := buf.String()

	// then
	if actual != expected {
		t.Errorf("For paths %v the expected pathString is %s, got %s", emptyPaths, expected, actual)
	}
}

func TestPathRendererTurnsValidSinglePathIntoValidPathString(t *testing.T) {
	// given
	segments := segments.Segments{
		segments.NewSegment(0, 1, 000),
		segments.NewSegment(1, 2, 100), // connectionID=1
		segments.NewSegment(2, 3, 200), // connectionID=2
		segments.NewSegment(3, 4, 300), // connectionID=3
	}
	path := pathfinding.Path{
		pathfinding.ConnectionID(1),
		pathfinding.ConnectionID(2),
		pathfinding.ConnectionID(3),
	}

	expected := "1-(100)-2-(200)-3-(300)-4"
	airportRenderer := stubAirportRenderer{}
	carrierRenderer := stubCarrierRenderer{}
	buf := bytes.NewBuffer([]byte{})
	pathRenderer := astext.NewPathRenderer(&airportRenderer, &carrierRenderer, segments, ",")

	// when
	pathRenderer.Render(buf, []pathfinding.Path{path})
	actual := buf.String()

	// then
	if actual != expected {
		t.Errorf("For path %v the expected pathString is %s, got %s", path, expected, actual)
	}
}

func TestPathRendererTurnsValidMultiplePathsIntoValidPathString(t *testing.T) {
	// given
	segments := segments.Segments{
		segments.NewSegment(0, 1, 000),
		segments.NewSegment(1, 2, 100), // connectionID=1
		segments.NewSegment(2, 3, 200), // connectionID=2
		segments.NewSegment(3, 4, 300), // connectionID=3
		segments.NewSegment(4, 1, 400), // connectionID=4
	}
	path1 := pathfinding.Path{
		pathfinding.ConnectionID(1),
		pathfinding.ConnectionID(2),
		pathfinding.ConnectionID(3),
	}
	path2 := pathfinding.Path{
		pathfinding.ConnectionID(4),
		pathfinding.ConnectionID(1),
		pathfinding.ConnectionID(2),
	}
	paths := []pathfinding.Path{path1, path2}
	expected := "1-(100)-2-(200)-3-(300)-4,4-(400)-1-(100)-2-(200)-3"
	airportRenderer := stubAirportRenderer{}
	carrierRenderer := stubCarrierRenderer{}
	buf := bytes.NewBuffer([]byte{})
	pathRenderer := astext.NewPathRenderer(&airportRenderer, &carrierRenderer, segments, ",")

	// when
	pathRenderer.Render(buf, paths)
	actual := buf.String()

	// then
	if actual != expected {
		t.Errorf("For path %v the expected pathString is %s, got %s", paths, expected, actual)
	}
}
