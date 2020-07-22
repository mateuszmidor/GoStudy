package util

import (
	"compress/gzip"
	"dataloading"
	"dataloading/csv"
	"fmt"
	"os"
)

// StartLoadingSegmentsFromGzipCSV begins loading raw segments from gzipped csv into output channel
func StartLoadingSegmentsFromGzipCSV(segmentsGzipCSV string, outSegments chan<- dataloading.RawSegment) {
	inputFile, err := os.Open(segmentsGzipCSV)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", segmentsGzipCSV, err)
		close(outSegments)
		return
	}
	defer inputFile.Close()

	gzipReader, err := gzip.NewReader(inputFile)
	if err != nil {
		fmt.Printf("Error creating GZIP reader %s: %v\n", segmentsGzipCSV, err)
		close(outSegments)
		return
	}
	defer gzipReader.Close()

	var loader csv.SegmentLoader
	loader.StartLoading(gzipReader, outSegments)
}

// StartLoadingAirportsFromGzipCSV begins loading raw airports from gzipped csv into output channel
func StartLoadingAirportsFromGzipCSV(airportsGzipCSV string, outAirports chan<- dataloading.RawAirport) {
	inputFile, err := os.Open(airportsGzipCSV)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", airportsGzipCSV, err)
		close(outAirports)
		return
	}
	defer inputFile.Close()

	gzipReader, err := gzip.NewReader(inputFile)
	if err != nil {
		fmt.Printf("Error creating GZIP reader %s: %v\n", airportsGzipCSV, err)
		close(outAirports)
		return
	}
	defer gzipReader.Close()

	var loader csv.AirportLoader
	loader.StartLoading(gzipReader, outAirports)
}

// StartLoadingNationsFromGzipCSV begins loading raw nations from gzipped csv into output channel
func StartLoadingNationsFromGzipCSV(nationsGzipCSV string, outNations chan<- dataloading.RawNation) {
	inputFile, err := os.Open(nationsGzipCSV)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", nationsGzipCSV, err)
		close(outNations)
		return
	}
	defer inputFile.Close()

	gzipReader, err := gzip.NewReader(inputFile)
	if err != nil {
		fmt.Printf("Error creating GZIP reader %s: %v\n", nationsGzipCSV, err)
		close(outNations)
		return
	}
	defer gzipReader.Close()

	var loader csv.NationsLoader
	loader.StartLoading(gzipReader, outNations)
}
