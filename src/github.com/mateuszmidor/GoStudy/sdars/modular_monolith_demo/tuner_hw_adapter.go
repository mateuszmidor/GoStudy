package main

import (
	"hexagons/hw"
	hwports "hexagons/hw/infrastructure"
	"hexagons/tuner"
	tunerports "hexagons/tuner/infrastructure"
)

// Implements tuner output ports towards hw, and hw output ports towards tuner
type TunerHwAdapter struct {
	tunerServicePort tunerports.ServicePort
	hwServicePort    hwports.ServicePort
}

func NewTunerHwAdapter(tuner *tuner.TunerRoot, hw *hw.HwRoot) TunerHwAdapter {
	adapter := TunerHwAdapter{tuner.GetServicePort(), hw.GetServicePort()}
	return adapter
}

// Hw -> Tuner
func (adapter *TunerHwAdapter) UpdateStationList(stationList []string) {
	adapter.tunerServicePort.StationListUpdated(stationList)
}

// Hw -> Tuner
func (adapter *TunerHwAdapter) UpdateSubscription(subscription bool) {
	adapter.tunerServicePort.SubscriptionUpdated(subscription)
}

// Tuner -> Hw
func (adapter *TunerHwAdapter) TuneToStation(stationId uint32) {
	adapter.hwServicePort.TuneToStation(stationId)
}
