package application

import (
	"fmt"
	"io"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/pathfinding"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/pathrendering/astext"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure"
)

type PathRendererAsText struct {
	writer io.Writer
}

func NewPathRendererAsText(w io.Writer) *PathRendererAsText {
	return &PathRendererAsText{writer: w}
}

func (r *PathRendererAsText) Render(paths []pathfinding.Path, flightsData *infrastructure.FlightsData) {
	airportRenderer := astext.NewLongAirportRenderer(flightsData.Airports)
	carrierRenderer := astext.NewShortCarrierRenderer(flightsData.Carriers)
	renderer := astext.NewPathRenderer(airportRenderer, carrierRenderer, flightsData.Segments, "\n")
	renderer.Render(r.writer, paths)

	// fmt.Fprint(r.writer, "\n")
	// fmt.Fprintf(r.writer, "[Total paths: %d, Took: %dms]", len(paths), elapsed.Milliseconds())
	fmt.Fprint(r.writer, "\n")
}
