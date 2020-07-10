package main

import (
	"airport"
	"carrier"
	"connections"
	"dataloading"
	"fmt"
	"pathfinding"
	"pathrendering"
	"segment"
)

type cliPathFinder struct {
	airports airport.Airports
	carriers carrier.Carriers
	segments segment.Segments
	network  connections.Network
}

func newCliPathFinder(segmentsGzipCSV string) *cliPathFinder {
	var rawSegments chan segment.RawSegment
	rawSegmentsLoader := dataloading.NewRawSegmentsFromCSVGzip(segmentsGzipCSV)

	// get airports used in segments
	rawSegments = make(chan segment.RawSegment, 100)
	go rawSegmentsLoader.StartLoadingSegments(rawSegments)
	airports := dataloading.NewRawSegmentsToAirportsFilter().Filter(rawSegments)

	// get carriers used in segments
	rawSegments = make(chan segment.RawSegment, 100)
	go rawSegmentsLoader.StartLoadingSegments(rawSegments)
	carriers := dataloading.NewRawSegmentsToCarriersFilter().Filter(rawSegments)

	// get actual segments
	rawSegments = make(chan segment.RawSegment, 100)
	go rawSegmentsLoader.StartLoadingSegments(rawSegments)
	segments := dataloading.NewRawSegmentsToSegmentsFilter(airports, carriers).Filter(rawSegments)

	// connection network
	network := connections.NewNetwork(segments)
	return &cliPathFinder{airports, carriers, segments, network}
}

func (f *cliPathFinder) findConnections(fromAirport, toAirport string) {
	from := pathfinding.NodeID(f.airports.GetByCode(fromAirport))
	to := pathfinding.NodeID(f.airports.GetByCode(toAirport))
	paths := pathfinding.FindPaths(from, to, &f.network)
	fmt.Println(f.pathsToString(paths))
	fmt.Println("Total paths count:", len(paths))
}

func (f *cliPathFinder) pathsToString(paths []pathfinding.Path) string {
	const pathSeparator = "\n"

	if len(paths) == 0 {
		return "<no paths found>"
	}

	airportRenderer := pathrendering.NewShortAirportRenderer(f.airports)
	carrierRenderer := pathrendering.NewShortCarrierRenderer(f.carriers)
	pathRenderer := pathrendering.NewRenderer(airportRenderer, carrierRenderer)

	var result string
	for _, p := range paths {
		result += pathRenderer.Render(p, f.segments) + pathSeparator
	}

	return result
}
