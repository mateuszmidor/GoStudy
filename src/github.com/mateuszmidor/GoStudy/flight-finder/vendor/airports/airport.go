package airports

import (
	"geo"
)

type ID int

const NullID = ID(-1)

type Airport struct {
	code      string // KRK
	name      string // Balice Krakow Airport
	longitude geo.Longitude
	latitude  geo.Latitude
}

func NewAirport(code, name string, longitude geo.Longitude, latitude geo.Latitude) Airport {
	return Airport{code, name, longitude, latitude}
}

func (a *Airport) Code() string {
	return a.code
}

func (a *Airport) Name() string {
	return a.name
}

func (a *Airport) SetName(name string) {
	a.name = name
}

func (a *Airport) SetCoordinates(lon geo.Longitude, lat geo.Latitude) {
	a.longitude = lon
	a.latitude = lat
}
