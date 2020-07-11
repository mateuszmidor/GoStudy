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
	fsegments, err := os.Open(segmentsGzipCSV)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", segmentsGzipCSV, err)
		close(outSegments)
		return
	}
	defer fsegments.Close()

	gzipReader, err := gzip.NewReader(fsegments)
	if err != nil {
		fmt.Printf("Error createing GZIP reader %s: %v\n", segmentsGzipCSV, err)
		close(outSegments)
		return
	}
	defer gzipReader.Close()

	var loader csv.SegmentLoader
	loader.StartLoading(gzipReader, outSegments)
}
