package util

import (
	"airports"
	"carriers"
	"connections"
	"dataloading"
	"errors"
	"fmt"
	"io"
	"pathfinding"
	"pathrendering/asjson"
	"pathrendering/astext"
	"segments"
	"sort"
	"time"
)

type ConnectionFinder struct {
	airports        airports.Airports
	carriers        carriers.Carriers
	segments        segments.Segments
	connections     pathfinding.Connections
	resultSeparator string
}

func NewConnectionFinder(segmentsGzipCSV, airportsGzipCSV string, resultSeparator string) *ConnectionFinder {
	var rawSegments chan dataloading.RawSegment

	// get airports used in segments
	rawSegments = make(chan dataloading.RawSegment, 100)
	go StartLoadingSegmentsFromGzipCSV(segmentsGzipCSV, rawSegments)
	airports := dataloading.FilterAirports(rawSegments)

	// get carriers used in segments
	rawSegments = make(chan dataloading.RawSegment, 100)
	go StartLoadingSegmentsFromGzipCSV(segmentsGzipCSV, rawSegments)
	carriers := dataloading.FilterCarriers(rawSegments)

	// get actual segments
	rawSegments = make(chan dataloading.RawSegment, 100)
	go StartLoadingSegmentsFromGzipCSV(segmentsGzipCSV, rawSegments)
	segments := dataloading.NewRawSegmentsToSegmentsFilter(airports, carriers).Filter(rawSegments)

	// enhance airports with name and location
	rawAirports := make(chan dataloading.RawAirport, 100)
	go StartLoadingAirportsFromGzipCSV(airportsGzipCSV, rawAirports)
	dataloading.EnrichAirports(airports, rawAirports)

	connections := connections.NewAdapter(segments)
	return &ConnectionFinder{airports, carriers, segments, connections, resultSeparator}
}

func (f *ConnectionFinder) FindConnectionsAsText(w io.Writer, fromAirport, toAirport string, maxSegmentCount int) {
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

	limiter := makeLimiter(maxSegmentCount)

	start := time.Now()
	paths := pathfinding.FindPaths(pathfinding.NodeID(from), pathfinding.NodeID(to), f.connections, limiter)
	elapsed := time.Now().Sub(start)

	f.pathsToText(w, paths)
	fmt.Fprint(w, f.resultSeparator)
	fmt.Fprintf(w, "[Total paths: %d, Took: %dms]", len(paths), elapsed.Milliseconds())
	fmt.Fprint(w, f.resultSeparator)
}

func (f *ConnectionFinder) pathsToText(w io.Writer, paths []pathfinding.Path) {
	airportRenderer := astext.NewLongAirportRenderer(f.airports)
	carrierRenderer := astext.NewShortCarrierRenderer(f.carriers)
	renderer := astext.NewPathRenderer(airportRenderer, carrierRenderer, f.segments, f.resultSeparator)
	renderer.Render(w, paths)
}

func (f *ConnectionFinder) FindConnectionsAsJSON(w io.Writer, fromAirport, toAirport string, maxSegmentCount int) error {
	from := f.airports.GetByCode(fromAirport)
	if from == airports.NullID {
		return errors.New("Invalid origin airport: " + fromAirport)
	}

	to := f.airports.GetByCode(toAirport)
	if to == airports.NullID {
		return errors.New("Invalid destination airport: " + toAirport)
	}

	// start := time.Now()
	limiter := makeLimiter(maxSegmentCount)
	paths := pathfinding.FindPaths(pathfinding.NodeID(from), pathfinding.NodeID(to), f.connections, limiter)
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
	// elapsed := time.Now().Sub(start)

	f.pathsToJSON(w, paths)
	// fmt.Fprint(w, f.pathsToString(paths))
	// fmt.Fprintf(w, "[Total paths: %d, Took: %dms]", len(paths), elapsed.Milliseconds())
	// fmt.Fprintln(w, f.resultSeparator)
	return nil
}

func (f *ConnectionFinder) pathsToJSON(w io.Writer, paths []pathfinding.Path) {
	renderer := asjson.NewPathRenderer(f.airports, f.carriers, f.segments)
	renderer.Render(w, paths)
}

func makeLimiter(maxSegmentCount int) pathfinding.CheckContinueBuildingPaths {
	return func(currentPathLen, totalPathsFound int) bool {
		maxPathLen := maxSegmentCount + 1 // KRK-WAW-GDN is 2 segments made of 3 airports
		return currentPathLen < maxPathLen && totalPathsFound < 1000
	}
}
