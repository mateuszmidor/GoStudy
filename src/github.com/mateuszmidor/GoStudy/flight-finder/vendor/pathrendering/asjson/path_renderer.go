package asjson

import (
	"airports"
	"carriers"
	"encoding/json"
	"io"
	"pathfinding"
	"segments"
)

// PathRenderer encodes slice of paths into JSON
type PathRenderer struct {
	airports airports.Airports
	carriers carriers.Carriers
	segments segments.Segments
}

type pathAdapter struct {
	FromAirport airportMarshaller   `json:"from_airport"`
	Segments    []segmentMarshaller `json:"segments"`
}

type airportMarshaller struct {
	AirportID airports.ID
	renderer  *PathRenderer
}

type segmentMarshaller struct {
	SegmentID segments.ID
	renderer  *PathRenderer
}

func (a *airportMarshaller) MarshalJSON() ([]byte, error) {
	view := NewJSONAirportView(&a.renderer.airports[a.AirportID])
	return json.Marshal(view)
}

func (s *segmentMarshaller) MarshalJSON() ([]byte, error) {
	seg := &s.renderer.segments[s.SegmentID]
	view := NewJSONSegmentView(&s.renderer.carriers[seg.Carrier()], &s.renderer.airports[seg.To()])
	return json.Marshal(view)
}

func NewPathRenderer(airports airports.Airports, carriers carriers.Carriers, segments segments.Segments) *PathRenderer {
	return &PathRenderer{airports: airports, carriers: carriers, segments: segments}
}

// Render serializes a slice of paths into JSON and writes it to "w"
func (r *PathRenderer) Render(w io.Writer, paths []pathfinding.Path) {
	adaptedPaths := r.adaptPaths(paths)
	json.NewEncoder(w).Encode(adaptedPaths)
}

func (r *PathRenderer) adaptPaths(paths []pathfinding.Path) []*pathAdapter {
	result := make([]*pathAdapter, 0, len(paths))
	for i := range paths {
		if adaptedPath := r.adaptPath(paths[i]); adaptedPath != nil {
			result = append(result, adaptedPath)
		}
	}

	return result
}

func (r *PathRenderer) adaptPath(path pathfinding.Path) *pathAdapter {
	if len(path) == 0 {
		return nil
	}

	var result pathAdapter
	s0 := &r.segments[path[0]]
	result.FromAirport = airportMarshaller{AirportID: s0.From(), renderer: r}
	result.Segments = make([]segmentMarshaller, len(path))

	for i, sID := range path {
		result.Segments[i] = segmentMarshaller{SegmentID: segments.ID(sID), renderer: r}
	}

	return &result
}
