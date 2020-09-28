package util

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"time"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/airports"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/carriers"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/connections"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/dataloading"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/nations"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/pathfinding"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/pathrendering/asjson"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/pathrendering/astext"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/segments"
)

type ConnectionFinder struct {
	nations         nations.Nations
	airports        airports.Airports
	carriers        carriers.Carriers
	segments        segments.Segments
	connections     pathfinding.Connections
	resultSeparator string
}

func NewConnectionFinder(segmentsGzipCSV, airportsGzipCSV, nationsGzipCSV string, resultSeparator string) *ConnectionFinder {
	var rawSegments chan dataloading.RawSegment

	// get nations
	rawNations := make(chan dataloading.RawNation, 100)
	go StartLoadingNationsFromGzipCSV(nationsGzipCSV, rawNations)
	nations := dataloading.FilterRawNations(rawNations)

	// get airports used in segments
	rawSegments = make(chan dataloading.RawSegment, 100)
	go StartLoadingSegmentsFromGzipCSV(segmentsGzipCSV, rawSegments)
	airportsUsedBySegments := dataloading.FilterAirports(rawSegments)

	// get carriers used in segments
	rawSegments = make(chan dataloading.RawSegment, 100)
	go StartLoadingSegmentsFromGzipCSV(segmentsGzipCSV, rawSegments)
	carriersUsedBySegments := dataloading.FilterCarriers(rawSegments)

	// get actual segments
	rawSegments = make(chan dataloading.RawSegment, 100)
	go StartLoadingSegmentsFromGzipCSV(segmentsGzipCSV, rawSegments)
	segments := dataloading.NewRawSegmentsToSegmentsFilter(airportsUsedBySegments, carriersUsedBySegments).Filter(rawSegments)

	// enhance airports with name and location
	rawAirports := make(chan dataloading.RawAirport, 100)
	go StartLoadingAirportsFromGzipCSV(airportsGzipCSV, rawAirports)
	dataloading.EnrichAirports(airportsUsedBySegments, rawAirports)

	connections := connections.NewAdapter(segments)
	return &ConnectionFinder{nations, airportsUsedBySegments, carriersUsedBySegments, segments, connections, resultSeparator}
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
	renderer := asjson.NewPathRenderer(f.airports, f.carriers, f.nations, f.segments)
	renderer.Render(w, paths)
}

func makeLimiter(maxSegmentCount int) pathfinding.CheckContinueBuildingPaths {
	return func(currentPathLen, totalPathsFound int) bool {
		maxPathLen := maxSegmentCount + 1 // KRK-WAW-GDN is 2 segments made of 3 airports
		return currentPathLen < maxPathLen && totalPathsFound < 1000
	}
}
