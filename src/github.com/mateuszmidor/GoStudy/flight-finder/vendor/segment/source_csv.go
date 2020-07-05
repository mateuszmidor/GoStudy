package segment

import (
	"encoding/csv"
	"io"
)

const numCSVColumns = 3

// SourceCSV loads segments from given source
type SourceCSV struct {
}

// StartLoadingSegments loads raw segments
// Pipeline instead batch load approach to accomodate segment database that would exceed machine ram limitations
// Usage: go source.StartLoadingSegments(...)
func (r *SourceCSV) StartLoadingSegments(reader io.Reader, outputSegments chan RawSegment) {
	csv := csv.NewReader(reader)
	csv.ReuseRecord = false
	csv.FieldsPerRecord = numCSVColumns

	for {
		rec, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err == nil && rec != nil {
			outputSegments <- parseRawSegment(rec)
		}

	}

	close(outputSegments)
}

func parseRawSegment(data []string) RawSegment {
	// CSV structure:
	// "KRK","KTW","LO"

	return RawSegment{
		FromAirportCode: data[0],
		ToAirportCode:   data[1],
		CarrierCode:     data[2],
	}
}
