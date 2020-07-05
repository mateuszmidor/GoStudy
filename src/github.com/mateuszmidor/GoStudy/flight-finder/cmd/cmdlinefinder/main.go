package main

import (
	"airport"
	"compress/gzip"
	"connections"
	"fmt"
	"multipathastar"
	"os"
	"runtime"
	"runtime/pprof"
	"segment"
)

const segments = "../../segments.csv.gz"
const segmentSeparator = "-"
const pathSeparator = "\n"

func main() {
	// collect CPU profile
	cpu, _ := os.Create("cpu.out")
	defer cpu.Close()
	pprof.StartCPUProfile(cpu)
	defer pprof.StopCPUProfile()

	findConnectionsDemo()

	// collect memory profile
	heap, _ := os.Create("mem.out")
	defer heap.Close()
	runtime.GC() // get up-to-date statistics
	pprof.WriteHeapProfile(heap)
}

func findConnectionsDemo() {
	airports := loadAirports()
	segments := loadSegments(airports)
	network := connections.NewNetwork(airports, segments)

	krk := multipathastar.NodeID(airports.GetByCode("WAW"))
	gdn := multipathastar.NodeID(airports.GetByCode("SEZ"))
	paths := multipathastar.FindPaths(krk, gdn, &network)
	fmt.Println(pathsToString(paths, network))
}

func loadAirports() airport.Airports {
	uniqueCodes := make(map[string]bool)
	segments := make(chan segment.RawSegment, 1)

	go startLoadingSegments(segments)
	for s := range segments {
		uniqueCodes[s.FromAirportCode] = true
		uniqueCodes[s.ToAirportCode] = true
	}

	var ab airport.Builder
	for code := range uniqueCodes {
		ab.Append(code, "")
	}
	return ab.Build()
}

func loadSegments(airports airport.Airports) segment.Segments {
	sb := segment.NewBuilder(airports)
	segments := make(chan segment.RawSegment, 100)

	go startLoadingSegments(segments)
	for s := range segments {
		sb.Append(s.FromAirportCode, s.ToAirportCode)
	}

	return sb.Build()
}

func startLoadingSegments(outSegments chan segment.RawSegment) {
	fsegments, err := os.Open(segments)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", segments, err)
		os.Exit(1)
	}
	defer fsegments.Close()

	gzipReader, err := gzip.NewReader(fsegments)
	if err != nil {
		fmt.Printf("Error createing GZIP reader %s: %v\n", segments, err)
		os.Exit(1)
	}
	defer gzipReader.Close()

	var source segment.SourceCSV
	source.StartLoadingSegments(gzipReader, outSegments)
}

func pathsToString(paths []multipathastar.Path, net connections.Network) string {
	if len(paths) == 0 {
		return ""
	}

	var result string
	for _, p := range paths {
		result += pathToString(p, net) + pathSeparator
	}

	return result
}

func pathToString(path multipathastar.Path, net connections.Network) string {
	if len(path) == 0 {
		return "<empty path>"
	}

	segment0 := net.GetSegment(segment.ID(path[0]))
	result := net.GetAirport(segment0.From()).Code() + segmentSeparator +
		// carrier
		net.GetAirport(segment0.To()).Code()
	// result := connections.nodes[connections.connections[path[0]].from].label + "-" +
	// 	connections.connections[path[0]].label + "-" +
	// 	connections.nodes[connections.connections[path[0]].to].label

	for i := 1; i < len(path); i++ {
		segment := net.GetSegment(segment.ID(path[i]))
		result += segmentSeparator + net.GetAirport(segment.To()).Code()
		// result += "-" + connections.connections[path[i]].label + "-" + connections.nodes[connections.connections[path[i]].to].label
	}

	return result
}
