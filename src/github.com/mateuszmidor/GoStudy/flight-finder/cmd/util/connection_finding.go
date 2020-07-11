package util

import (
	"airport"
	"carrier"
	"connections"
	"dataloading"
	"fmt"
	"io"
	"pathfinding"
	"pathrendering"
	"segment"
	"time"
)

type ConnectionFinder struct {
	airports        airport.Airports
	carriers        carrier.Carriers
	segments        segment.Segments
	connections     connections.Connections
	resultSeparator string
}

func NewConnectionFinder(segmentsGzipCSV string, resultSeparator string) *ConnectionFinder {
	var rawSegments chan dataloading.RawSegment

	// get airports used in segments
	rawSegments = make(chan dataloading.RawSegment, 100)
	go StartLoadingSegmentsFromGzipCSV(segmentsGzipCSV, rawSegments)
	airports := dataloading.NewRawSegmentsToAirportsFilter().Filter(rawSegments)

	// get carriers used in segments
	rawSegments = make(chan dataloading.RawSegment, 100)
	go StartLoadingSegmentsFromGzipCSV(segmentsGzipCSV, rawSegments)
	carriers := dataloading.NewRawSegmentsToCarriersFilter().Filter(rawSegments)

	// get actual segments
	rawSegments = make(chan dataloading.RawSegment, 100)
	go StartLoadingSegmentsFromGzipCSV(segmentsGzipCSV, rawSegments)
	segments := dataloading.NewRawSegmentsToSegmentsFilter(airports, carriers).Filter(rawSegments)

	// connection connections
	connections := connections.NewConnections(segments)
	return &ConnectionFinder{airports, carriers, segments, connections, resultSeparator}
}

func (f *ConnectionFinder) FindConnections(w io.Writer, fromAirport, toAirport string) {
	start := time.Now()
	from := pathfinding.NodeID(f.airports.GetByCode(fromAirport))
	to := pathfinding.NodeID(f.airports.GetByCode(toAirport))
	paths := pathfinding.FindPaths(from, to, &f.connections)
	d := time.Now().Sub(start)

	fmt.Fprint(w, f.pathsToString(paths))
	fmt.Fprintf(w, "[Total paths: %d, Took: %dms]", len(paths), d.Milliseconds())
	fmt.Fprintln(w, f.resultSeparator)
}

func (f *ConnectionFinder) pathsToString(paths []pathfinding.Path) string {

	if len(paths) == 0 {
		return "<no paths found>"
	}

	airportRenderer := pathrendering.NewShortAirportRenderer(f.airports)
	carrierRenderer := pathrendering.NewShortCarrierRenderer(f.carriers)
	pathRenderer := pathrendering.NewRenderer(airportRenderer, carrierRenderer)

	var result string
	for _, p := range paths {
		result += pathRenderer.Render(p, f.segments) + f.resultSeparator
	}

	return result
}
