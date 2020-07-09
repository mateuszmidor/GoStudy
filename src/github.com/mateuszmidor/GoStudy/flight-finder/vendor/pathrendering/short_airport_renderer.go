package pathrendering

import (
	"airport"
)

// ShortAirportRenderer renders short string representation of airportID
type ShortAirportRenderer struct {
	airports airport.Airports
}

// NewShortAirportRenderer is constructor
func NewShortAirportRenderer(airports airport.Airports) *ShortAirportRenderer {
	return &ShortAirportRenderer{airports}
}

// Render creates string representation of airport
func (r *ShortAirportRenderer) Render(id airport.AirportID) string {
	return r.airports[id].Code()
}
