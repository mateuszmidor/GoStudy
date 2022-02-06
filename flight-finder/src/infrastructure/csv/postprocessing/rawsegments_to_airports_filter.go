package postprocessing

import (
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/airports"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure/csv/loading"
)

func ExtractAirports(segments <-chan loading.CSVSegment) airports.Airports {
	uniqueCodes := make(map[string]bool)

	for s := range segments {
		uniqueCodes[s.FromAirportCode] = true
		uniqueCodes[s.ToAirportCode] = true
	}

	var ab airports.Builder
	for code := range uniqueCodes {
		ab.Append(code)
	}

	return ab.Build()
}
