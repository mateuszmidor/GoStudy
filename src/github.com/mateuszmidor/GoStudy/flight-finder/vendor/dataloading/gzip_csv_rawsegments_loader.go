package dataloading

import (
	"compress/gzip"
	"fmt"
	"os"
	"segment"
)

type GzipCSVRawSegmentsLoader struct {
	gzipFilename string
}

func NewRawSegmentsFromCSVGzip(filename string) *GzipCSVRawSegmentsLoader {
	return &GzipCSVRawSegmentsLoader{filename}
}

func (r *GzipCSVRawSegmentsLoader) StartLoadingSegments(outSegments chan segment.RawSegment) {
	fsegments, err := os.Open(r.gzipFilename)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", r.gzipFilename, err)
		return
	}
	defer fsegments.Close()

	gzipReader, err := gzip.NewReader(fsegments)
	if err != nil {
		fmt.Printf("Error createing GZIP reader %s: %v\n", r.gzipFilename, err)
		return
	}
	defer gzipReader.Close()

	var source SourceCSV
	source.StartLoadingSegments(gzipReader, outSegments)
}
