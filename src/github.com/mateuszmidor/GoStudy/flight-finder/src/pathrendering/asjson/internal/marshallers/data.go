package marshallers

import (
	"github.com/mateuszmidor/GoStudy/flight-finder/src/airports"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/carriers"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/nations"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/segments"
)

// Data holds details of airports/carriers/segments
type Data struct {
	Airports airports.Airports
	Carriers carriers.Carriers
	Nations  nations.Nations
	Segments segments.Segments
}
