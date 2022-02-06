package postprocessing

import (
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/airports"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure/csv/loading"
)

func EnrichAirports(poorAirports airports.Airports, csvAirports <-chan loading.CSVAirport) {
	for ra := range csvAirports {
		id := poorAirports.GetByCode(ra.AirportCode)
		if id == airports.NullID {
			continue
		}
		poorAirports[id].SetName(ra.FullName)
		poorAirports[id].SetNation(ra.Nation)
		poorAirports[id].SetCoordinates(ra.Longitude, ra.Latitude)
	}
}
