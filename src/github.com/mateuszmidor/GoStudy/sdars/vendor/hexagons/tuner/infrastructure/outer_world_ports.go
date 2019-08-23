package infrastructure

import "hexagons/tuner/domain"

// HwPort allows Tuner talk to Hardware
type HwPort interface {
	TuneToStation(stationID domain.StationID)
}

// UIPort allows Tuner talk to UI
type UIPort interface {
	UpdateStationList(stationList domain.StationList)
	UpdateSubscription(subscription domain.Subscription)
}

// OuterWorldPorts collects the ports that Tuner can talk to
type OuterWorldPorts struct {
	UIPort UIPort
	HwPort HwPort
}
