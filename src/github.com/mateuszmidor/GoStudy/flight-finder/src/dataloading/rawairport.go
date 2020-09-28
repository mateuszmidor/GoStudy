package dataloading

import "github.com/mateuszmidor/GoStudy/flight-finder/src/geo"

// RawAirport is airport in its denormalized form
type RawAirport struct {
	AirportCode string        // eg. "GDN"
	FullName    string        // eg. "Gdynia Airport"
	Nation      string        // eg. "PL"
	Longitude   geo.Longitude // eg. 20.1
	Latitude    geo.Latitude  // eg. -50.4
}

// NewRawAirport is constructor
func NewRawAirport(code, name, nation string, lng geo.Longitude, lat geo.Latitude) RawAirport {
	return RawAirport{
		AirportCode: code,
		FullName:    name,
		Nation:      nation,
		Longitude:   lng,
		Latitude:    lat,
	}
}
