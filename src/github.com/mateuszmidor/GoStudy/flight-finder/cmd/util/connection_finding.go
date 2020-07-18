package util

import (
	"airports"
	"carriers"
	"connections"
	"dataloading"
	"fmt"
	"io"
	"pathfinding"
	"pathrendering/asjson"
	"pathrendering/astext"
	"segments"
	"strings"
	"time"
)

type ConnectionFinder struct {
	airports        airports.Airports
	carriers        carriers.Carriers
	segments        segments.Segments
	connections     pathfinding.Connections
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

	// enhance airports with name and location
	rawAirports := make(chan dataloading.RawAirport, 100)
	go StartLoadingAirportsFromGzipCSV("../../airports.csv.gz", rawAirports)
	dataloading.EnrichAirports(airports, rawAirports)

	connections := connections.NewAdapter(segments)
	return &ConnectionFinder{airports, carriers, segments, &connections, resultSeparator}
}

func (f *ConnectionFinder) FindConnectionsAsText(w io.Writer, fromAirport, toAirport string) {
	from := f.airports.GetByCode(fromAirport)
	if from == airports.NullID {
		fmt.Fprintf(w, "Invalid from airport: %s%s", fromAirport, f.resultSeparator)
		return
	}

	to := f.airports.GetByCode(toAirport)
	if to == airports.NullID {
		fmt.Fprintf(w, "Invalid to airport: %s%s", toAirport, f.resultSeparator)
		return
	}

	start := time.Now()
	paths := pathfinding.FindPaths(pathfinding.NodeID(from), pathfinding.NodeID(to), f.connections)
	elapsed := time.Now().Sub(start)

	fmt.Fprint(w, f.pathsToString(paths))
	fmt.Fprintf(w, "[Total paths: %d, Took: %dms]", len(paths), elapsed.Milliseconds())
	fmt.Fprintln(w, f.resultSeparator)
}

func (f *ConnectionFinder) FindConnectionsAsJSON(w io.Writer, fromAirport, toAirport string) {
	from := f.airports.GetByCode(fromAirport)
	if from == airports.NullID {
		fmt.Fprintf(w, "Invalid from airport: %s%s", fromAirport, f.resultSeparator)
		return
	}

	to := f.airports.GetByCode(toAirport)
	if to == airports.NullID {
		fmt.Fprintf(w, "Invalid to airport: %s%s", toAirport, f.resultSeparator)
		return
	}

	// start := time.Now()
	paths := pathfinding.FindPaths(pathfinding.NodeID(from), pathfinding.NodeID(to), f.connections)
	// elapsed := time.Now().Sub(start)

	f.pathsToJSON(w, paths)
	// fmt.Fprint(w, f.pathsToString(paths))
	// fmt.Fprintf(w, "[Total paths: %d, Took: %dms]", len(paths), elapsed.Milliseconds())
	// fmt.Fprintln(w, f.resultSeparator)
}

func (f *ConnectionFinder) pathsToString(paths []pathfinding.Path) string {

	if len(paths) == 0 {
		return "<no paths found>"
	}

	airportRenderer := astext.NewLongAirportRenderer(f.airports)
	carrierRenderer := astext.NewShortCarrierRenderer(f.carriers)
	pathRenderer := astext.NewPathRenderer(airportRenderer, carrierRenderer)

	var sb strings.Builder
	for i := range paths {
		sb.WriteString(pathRenderer.Render(paths[i], f.segments))
		sb.WriteString(f.resultSeparator)
	}

	return sb.String()
}

func (f *ConnectionFinder) pathsToJSON(w io.Writer, paths []pathfinding.Path) {
	renderer := asjson.NewPathRenderer(f.airports, f.carriers, f.segments)
	renderer.Render(w, paths)
}
