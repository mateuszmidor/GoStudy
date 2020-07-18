package asjson

import (
	"airports"
	"carriers"
	"encoding/json"
	"io"
	"pathfinding"
	"pathrendering/asjson/internal/marshallers"
	"segments"
)

// PathRenderer encodes []pathfinding.Path into JSON
type PathRenderer struct {
	data marshallers.Data
}

// NewPathRenderer is constructor
func NewPathRenderer(airports airports.Airports, carriers carriers.Carriers, segments segments.Segments) *PathRenderer {
	return &PathRenderer{
		data: marshallers.Data{
			Airports: airports,
			Carriers: carriers,
			Segments: segments,
		},
	}
}

// Render encodes []pathfinding.Path into JSON
func (r *PathRenderer) Render(w io.Writer, paths []pathfinding.Path) {
	pathsMarshaller := buildMarshallerForPaths(paths, &r.data)
	json.NewEncoder(w).Encode(pathsMarshaller)
}
