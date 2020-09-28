package marshallers

import (
	"airports"
	"carriers"
	"nations"
	"segments"
)

// Data holds details of airports/carriers/segments
type Data struct {
	Airports airports.Airports
	Carriers carriers.Carriers
	Nations  nations.Nations
	Segments segments.Segments
}
