package astext_test

import (
	"airports"
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

func TestPathRendererTurnsValidPathIntoValidPathString(t *testing.T) {
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

	// when
	airportRenderer := stubAirportRenderer{}
	carrierRenderer := stubCarrierRenderer{}
	pathRenderer := astext.NewPathRenderer(&airportRenderer, &carrierRenderer)
	actual := pathRenderer.Render(path, segments)

	// then
	if actual != expected {
		t.Errorf("For path %v the expected pathString is %s, got %s", path, expected, actual)
	}
}

func TestPathRendererTurnsEmptyPathIntoEmptyPathString(t *testing.T) {
	// given
	segments := segments.Segments{}
	path := pathfinding.Path{}

	expected := "<empty path>"

	// when
	airportRenderer := stubAirportRenderer{}
	carrierRenderer := stubCarrierRenderer{}
	pathRenderer := astext.NewPathRenderer(&airportRenderer, &carrierRenderer)
	actual := pathRenderer.Render(path, segments)

	// then
	if actual != expected {
		t.Errorf("For path %v the expected pathString is %s, got %s", path, expected, actual)
	}
}
