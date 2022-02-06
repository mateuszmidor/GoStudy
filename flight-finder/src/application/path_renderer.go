package application

import (
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/pathfinding"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure"
)

type PathRenderer interface {
	Render(paths []pathfinding.Path, flightsData *infrastructure.FlightsData)
}
