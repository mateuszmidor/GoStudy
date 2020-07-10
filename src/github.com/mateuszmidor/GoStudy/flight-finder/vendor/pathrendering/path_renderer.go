package pathrendering

import (
	"airport"
	"carrier"
	"fmt"
	"pathfinding"
	"segment"
)

type AirportRenderer interface {
	Render(airport.ID) string
}

type CarrierRenderer interface {
	Render(carrier.ID) string
}

type PathRenderer struct {
	airportRenderer AirportRenderer
	carrierRenderer CarrierRenderer
}

func NewRenderer(airportRenderer AirportRenderer, carrierRenderer CarrierRenderer) *PathRenderer {
	return &PathRenderer{airportRenderer, carrierRenderer}
}

func (r *PathRenderer) Render(path pathfinding.Path, segments segment.Segments) string {
	if len(path) == 0 {
		return "<empty path>"
	}

	s0 := segments[path[0]]
	result := r.airportRenderer.Render(s0.From())

	for _, sID := range path {
		s := segments[sID]
		result += fmt.Sprintf("-(%s)-%s", r.carrierRenderer.Render(s.Carrier()), r.airportRenderer.Render(s.To()))
	}

	return result
}
