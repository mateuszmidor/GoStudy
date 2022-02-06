package asjson

import (
	"encoding/json"
	"io"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/airports"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/carriers"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/nations"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/pathfinding"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/pathrendering/asjson/internal/marshallers"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/segments"
)

// PathRenderer encodes []pathfinding.Path into JSON
type PathRenderer struct {
	data marshallers.Data
}

// NewPathRenderer is constructor
func NewPathRenderer(airports airports.Airports, carriers carriers.Carriers, nations nations.Nations, segments segments.Segments) *PathRenderer {
	return &PathRenderer{
		data: marshallers.Data{
			Airports: airports,
			Carriers: carriers,
			Nations:  nations,
			Segments: segments,
		},
	}
}

// Render encodes []pathfinding.Path into JSON
func (r *PathRenderer) Render(w io.Writer, paths []pathfinding.Path) {
	pathsMarshaller := buildMarshallerForPaths(paths, &r.data)
	json.NewEncoder(w).Encode(pathsMarshaller)
}
