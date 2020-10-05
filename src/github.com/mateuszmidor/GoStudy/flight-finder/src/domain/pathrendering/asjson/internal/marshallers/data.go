package marshallers

import (
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/airports"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/carriers"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/nations"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/segments"
)

// Data holds details of airports/carriers/segments
type Data struct {
	Airports airports.Airports
	Carriers carriers.Carriers
	Nations  nations.Nations
	Segments segments.Segments
}
