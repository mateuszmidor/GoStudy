package infrastructure

import "hexagons/tuner/domain"

// what Tuner can tell to UI
type UiPortOut interface {
	UpdateStationList(stationList domain.StationList)
	UpdateSubscription(subscription domain.Subscription)
}