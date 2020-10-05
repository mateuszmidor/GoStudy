package application

import (
	"fmt"
	"sort"

	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/airports"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/connections"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/domain/pathfinding"
	"github.com/mateuszmidor/GoStudy/flight-finder/src/infrastructure"
)

type ConnectionFindingService struct {
	flightsData infrastructure.FlightsData
	connections pathfinding.Connections
}

func NewConnectionFindingService(repo infrastructure.FlightsDataRepo) *ConnectionFindingService {

	flightsData := repo.Load()
	connections := connections.NewAdapter(flightsData.Segments)
	return &ConnectionFindingService{flightsData: flightsData, connections: connections}
}

func (f *ConnectionFindingService) Find(fromAirport, toAirport string, maxSegmentCount int, pathRenderer PathRenderer) error {
	flightsData := &f.flightsData
	from := flightsData.Airports.GetByCode(fromAirport)
	if from == airports.NullID {
		return fmt.Errorf("Invalid from airport: %s%s", fromAirport, "\n")
	}

	to := flightsData.Airports.GetByCode(toAirport)
	if to == airports.NullID {
		return fmt.Errorf("Invalid to airport: %s%s", toAirport, "\n")
	}

	limiter := makeLimiter(maxSegmentCount)

	// start := time.Now()
	paths := pathfinding.FindPaths(pathfinding.NodeID(from), pathfinding.NodeID(to), f.connections, limiter)
	// elapsed := time.Now().Sub(start)
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	pathRenderer.Render(paths, flightsData)
	return nil
}

func makeLimiter(maxSegmentCount int) pathfinding.CheckContinueBuildingPaths {
	return func(currentPathLen, totalPathsFound int) bool {
		maxPathLen := maxSegmentCount + 1 // KRK-WAW-GDN is 2 segments made of 3 airports
		return currentPathLen < maxPathLen && totalPathsFound < 1000
	}
}
