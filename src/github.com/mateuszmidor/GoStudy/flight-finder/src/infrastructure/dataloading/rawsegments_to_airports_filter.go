package dataloading

import "github.com/mateuszmidor/GoStudy/flight-finder/src/domain/airports"

func FilterAirports(segments <-chan RawSegment) airports.Airports {
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
