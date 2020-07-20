package astext

import (
	"airports"
	"carriers"
	"fmt"
	"io"
	"pathfinding"
	"segments"
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
	segments        segments.Segments
	separator       string
}

func NewPathRenderer(airportRenderer AirportRenderer, carrierRenderer CarrierRenderer, segments segments.Segments, separator string) *PathRenderer {
	return &PathRenderer{
		airportRenderer: airportRenderer,
		carrierRenderer: carrierRenderer,
		segments:        segments,
		separator:       separator,
	}
}

func (r *PathRenderer) Render(w io.Writer, paths []pathfinding.Path) {
	for i := range paths {
		if i > 0 {
			w.Write([]byte(r.separator))
		}
		r.renderSinglePath(w, paths[i])
	}
}

func (r *PathRenderer) renderSinglePath(w io.Writer, path pathfinding.Path) {
	if len(path) == 0 {
		return
	}

	s0 := r.segments[path[0]]
	w.Write([]byte(r.airportRenderer.Render(s0.From())))
	for _, sID := range path {
		s := r.segments[sID]
		fmt.Fprintf(w, "-(%s)-%s", r.carrierRenderer.Render(s.Carrier()), r.airportRenderer.Render(s.To()))
	}
}
