package dataloading

import "airports"

func EnrichAirports(poorAirports airports.Airports, rawAirports <-chan RawAirport) {
	for ra := range rawAirports {
		id := poorAirports.GetByCode(ra.AirportCode)
		if id == airports.NullID {
			continue
		}
		poorAirports[id].SetName(ra.FullName)
		poorAirports[id].SetNation(ra.Nation)
		poorAirports[id].SetCoordinates(ra.Longitude, ra.Latitude)
	}
}
