package infrastructure

import "hexagons/tuner/domain"

// what Tuner can tell to GUI
type GuiPortOut interface {
	UpdateStationList(stationList domain.StationList)
	UpdateSubscription(subscription domain.Subscription)
}