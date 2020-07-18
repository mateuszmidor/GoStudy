package astext

import (
	"airports"
	"carriers"
	"fmt"
	"pathfinding"
	"segments"
	"strings"
)

type AirportRenderer interface {
	Render(airports.ID) string
}

type CarrierRenderer interface {
	Render(carriers.ID) string
}

type PathRenderer struct {
	airportRenderer AirportRenderer
	carrierRenderer CarrierRenderer
}

func NewPathRenderer(airportRenderer AirportRenderer, carrierRenderer CarrierRenderer) *PathRenderer {
	return &PathRenderer{airportRenderer, carrierRenderer}
}

func (r *PathRenderer) Render(path pathfinding.Path, segments segments.Segments) string {
	if len(path) == 0 {
		return "<empty path>"
	}

	s0 := segments[path[0]]
	var sb strings.Builder
	sb.WriteString(r.airportRenderer.Render(s0.From()))
	for _, sID := range path {
		s := segments[sID]
		sb.WriteString(fmt.Sprintf("-(%s)-%s", r.carrierRenderer.Render(s.Carrier()), r.airportRenderer.Render(s.To())))
	}

	return sb.String()
}
