package views

import (
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/airports"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/nations"
)

// Airport is json view of airports.Airport
type Airport struct {
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Nation         string  `json:"nation"`
	NationFullName string  `json:"nation_full_name"`
	Lon            float32 `json:"lon"`
	Lat            float32 `json:"lat"`
}

// NewJSONAirportView is constructor
func NewJSONAirportView(a *airports.Airport, nation *nations.Nation) *Airport {
	nationFullName := "<unknown>"
	if nation != nil {
		nationFullName = nation.Name()
	}
	return &Airport{
		Code:           a.Code(),
		Name:           a.Name(),
		Nation:         a.Nation(),
		NationFullName: nationFullName,
		Lon:            float32(a.Longitude()),
		Lat:            float32(a.Latitude()),
	}
}
