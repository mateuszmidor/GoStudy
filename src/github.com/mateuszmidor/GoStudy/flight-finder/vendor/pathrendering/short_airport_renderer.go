package pathrendering

import (
	"airports"
)

// ShortAirportRenderer renders short string representation of ID
type ShortAirportRenderer struct {
	airports airports.Airports
}

// NewShortAirportRenderer is constructor
func NewShortAirportRenderer(airports airports.Airports) *ShortAirportRenderer {
	return &ShortAirportRenderer{airports}
}

// Render creates string representation of airport
func (r *ShortAirportRenderer) Render(id airports.ID) string {
	return r.airports[id].Code()
}
