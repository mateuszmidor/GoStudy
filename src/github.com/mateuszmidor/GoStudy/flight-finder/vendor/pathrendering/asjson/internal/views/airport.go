package views

import "airports"

// Airport is json view of airports.Airport
type Airport struct {
	Code string  `json:"code"`
	Name string  `json:"name"`
	Lon  float32 `json:"lon"`
	Lat  float32 `json:"lat"`
}

// NewJSONAirportView is constructor
func NewJSONAirportView(a *airports.Airport) *Airport {
	return &Airport{
		Code: a.Code(),
		Name: a.Name(),
		Lon:  float32(a.Longitude()),
		Lat:  float32(a.Latitude()),
	}
}
