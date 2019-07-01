package infrastructure

import "hexagons/tuner/domain"

// what GUI can tell to tuner
type GuiPortIn interface {
	TuneToStation(stationId domain.StationId)
}