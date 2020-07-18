package astext

import (
	"airports"
)

// LongAirportRenderer renders short string representation of ID
type LongAirportRenderer struct {
	airports airports.Airports
}

// NewLongAirportRenderer is constructor
func NewLongAirportRenderer(airports airports.Airports) *LongAirportRenderer {
	return &LongAirportRenderer{airports}
}

// Render creates string representation of airport
func (r *LongAirportRenderer) Render(id airports.ID) string {
	return r.airports[id].Name()
}
