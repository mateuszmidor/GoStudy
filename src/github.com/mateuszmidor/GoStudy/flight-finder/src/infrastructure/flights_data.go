package infrastructure

import (
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/airports"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/carriers"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/nations"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/segments"
)

// FlightsData is just a flights data pack
type FlightsData struct {
	Airports airports.Airports
	Carriers carriers.Carriers
	Nations  nations.Nations
	Segments segments.Segments
}
