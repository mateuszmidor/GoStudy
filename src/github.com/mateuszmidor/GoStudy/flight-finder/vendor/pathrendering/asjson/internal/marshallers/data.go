package marshallers

import (
	"airports"
	"carriers"
	"segments"
)

// Data holds details of airports/carriers/segments
type Data struct {
	Airports airports.Airports
	Carriers carriers.Carriers
	Segments segments.Segments
}
