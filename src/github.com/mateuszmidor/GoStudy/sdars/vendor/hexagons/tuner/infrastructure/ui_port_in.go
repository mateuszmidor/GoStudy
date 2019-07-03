package infrastructure

import "hexagons/tuner/domain"

// what UI can tell to tuner
type UiPortIn interface {
	TuneToStation(stationId domain.StationId)
}