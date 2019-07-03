package infrastructure

import "hexagons/tuner/domain"

// what Tuner can tell to Hardware
type HwPortOut interface {
	TuneToStation(stationId domain.StationId)
}