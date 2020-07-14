package dataloading

import (
	"geo"
)

// RawAirport is airport in its denormalized form
type RawAirport struct {
	AirportCode string        // eg. "GDN"
	FullName    string        // eg. "Gdynia Airport"
	Longitude   geo.Longitude // eg. 20.1
	Latitude    geo.Latitude  // eg. -50.4
}

// NewRawAirport is constructor
func NewRawAirport(code, name string, lng geo.Longitude, lat geo.Latitude) RawAirport {
	return RawAirport{
		AirportCode: code,
		FullName:    name,
		Longitude:   lng,
		Latitude:    lat,
	}
}
