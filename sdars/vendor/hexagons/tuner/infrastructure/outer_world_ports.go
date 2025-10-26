package infrastructure

import "sharedkernel"

// HwPort allows Tuner talk to Hardware
type HwPort interface {
	TuneToStation(stationID sharedkernel.StationID)
}

// UIPort allows Tuner talk to UI
type UIPort interface {
	UpdateStationList(stationList sharedkernel.StationList)
	UpdateSubscription(subscription sharedkernel.Subscription)
}

// OuterWorldPorts collects the ports that Tuner can talk to
type OuterWorldPorts struct {
	UIPort UIPort
	HwPort HwPort
}
