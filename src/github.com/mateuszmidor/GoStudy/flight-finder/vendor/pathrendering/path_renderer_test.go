package pathrendering_test

import (
	"airport"
	"carrier"
	"pathfinding"
	"pathrendering"
	"segment"
	"strconv"
	"testing"
)

type stubAirportRenderer struct {
}

func (r *stubAirportRenderer) Render(id airport.AirportID) string {
	return strconv.Itoa(int(id))
}

type stubCarrierRenderer struct {
}

func (r *stubCarrierRenderer) Render(id carrier.ID) string {
	return strconv.Itoa(int(id))
}

func TestPathRendererTurnsValidPathIntoValidPathString(t *testing.T) {
	// given
	segments := segment.Segments{
		segment.NewSegment(0, 1, 000),
		segment.NewSegment(1, 2, 100), // connectionID=1
		segment.NewSegment(2, 3, 200), // connectionID=2
		segment.NewSegment(3, 4, 300), // connectionID=3
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
	pathRenderer := pathrendering.NewRenderer(&airportRenderer, &carrierRenderer)
	actual := pathRenderer.Render(path, segments)

	// then
	if actual != expected {
		t.Errorf("For path %v the expected pathString is %s, got %s", path, expected, actual)
	}
}

func TestPathRendererTurnsEmptyPathIntoEmptyPathString(t *testing.T) {
	// given
	segments := segment.Segments{}
	path := pathfinding.Path{}

	expected := "<empty path>"

	// when
	airportRenderer := stubAirportRenderer{}
	carrierRenderer := stubCarrierRenderer{}
	pathRenderer := pathrendering.NewRenderer(&airportRenderer, &carrierRenderer)
	actual := pathRenderer.Render(path, segments)

	// then
	if actual != expected {
		t.Errorf("For path %v the expected pathString is %s, got %s", path, expected, actual)
	}
}
