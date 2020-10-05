package csv

import (
	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure/dataloading"
)

// FlightsDataRepoCSV implements FlightsDataRepo
type FlightsDataRepoCSV struct {
	csvFilesDirectory string
}

// NewFlightsDataRepoCSV is constructor
func NewFlightsDataRepoCSV(csvFilesDirectory string) *FlightsDataRepoCSV {
	return &FlightsDataRepoCSV{csvFilesDirectory: csvFilesDirectory}
}

// Load reads CSV files and returns flights data
func (r *FlightsDataRepoCSV) Load() infrastructure.FlightsData {
	airportsGzipCSV := r.csvFilesDirectory + "airports.csv.gz"
	nationsGzipCSV := r.csvFilesDirectory + "nations.csv.gz"
	segmentsGzipCSV := r.csvFilesDirectory + "segments.csv.gz"

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

	// enhance airports with name, nation and location
	rawAirports := make(chan dataloading.RawAirport, 100)
	go StartLoadingAirportsFromGzipCSV(airportsGzipCSV, rawAirports)
	dataloading.EnrichAirports(airportsUsedBySegments, rawAirports)

	return infrastructure.FlightsData{
		Nations:  nations,
		Airports: airportsUsedBySegments,
		Carriers: carriersUsedBySegments,
		Segments: segments,
	}

}
