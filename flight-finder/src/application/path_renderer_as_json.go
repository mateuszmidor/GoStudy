package application

import (
	"fmt"
	"io"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/pathfinding"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/pathrendering/asjson"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure"
)

type PathRendererAsJSON struct {
	writer io.Writer
}

func NewPathRendererAsJSON(w io.Writer) *PathRendererAsJSON {
	return &PathRendererAsJSON{writer: w}
}

func (r *PathRendererAsJSON) Render(paths []pathfinding.Path, flightsData *infrastructure.FlightsData) {
	renderer := asjson.NewPathRenderer(flightsData.Airports, flightsData.Carriers, flightsData.Nations, flightsData.Segments)
	renderer.Render(r.writer, paths)

	// fmt.Fprint(r.writer, "\n")
	// fmt.Fprintf(r.writer, "[Total paths: %d, Took: %dms]", len(paths), elapsed.Milliseconds())
	fmt.Fprint(r.writer, "\n")
}
