package infrastructure

import "hexagons/tuner/domain"

// what Tuner can tell to Hardware
type HardwarePortOut interface {
	TuneToStation(stationId domain.StationId)
}