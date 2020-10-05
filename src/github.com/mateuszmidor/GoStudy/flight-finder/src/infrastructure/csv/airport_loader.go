package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/geo"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure/dataloading"
)

const numAirportCSVColumns = 11

// AirportLoader loads airports from given source
type AirportLoader struct {
}

// StartLoading starts loading raw airports into output channel
// Pipeline instead batch load approach to accomodate segment database that would exceed machine ram limitations
// Usage: go source.StartLoading(...)
func (r *AirportLoader) StartLoading(reader io.Reader, outputAirports chan<- dataloading.RawAirport) {
	csv := csv.NewReader(reader)
	csv.ReuseRecord = true
	csv.FieldsPerRecord = numAirportCSVColumns

	for {
		rec, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err == nil && rec != nil {
			airport, err := parseRawAirport(rec)
			if err != nil {
				fmt.Printf("AiportLoader.StartLoading error: %v %+v\n", err.Error(), airport)
			}

			outputAirports <- airport
		}

	}

	close(outputAirports)
}

func parseRawAirport(data []string) (dataloading.RawAirport, error) {
	// CSV structure:
	// MARKET,LATDEG,LATMIN,LATSEC,LNGDEG,LNGMIN,LNGSEC,LATHEM,LNGHEM,NATION,DESCRIPTION
	var errorString string
	var coords [6]int
	for i := range coords {
		n, err := strconv.Atoi(data[i+1])
		if err != nil {
			errorString += "Error converting to int: " + data[i+1] + " "
		}
		coords[i] = n
	}

	latHem := data[7]
	lngHem := data[8]

	lat, err := geo.ConvertDegMinSecHemToLatitude(coords[0], coords[1], coords[2], latHem)
	if err != nil {
		errorString += err.Error() + " "
	}

	lng, err := geo.ConvertDegMinSecHemToLongitude(coords[3], coords[4], coords[5], lngHem)
	if err != nil {
		errorString += err.Error() + " "
	}

	if len(errorString) != 0 {
		return dataloading.NewRawAirport(data[0], data[10], data[9], lng, lat), errors.New(errorString)
	}
	return dataloading.NewRawAirport(data[0], data[10], data[9], lng, lat), nil
}
