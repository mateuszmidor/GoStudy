package main

import (
	"hexagons/hw"
	hwports "hexagons/hw/infrastructure"
	"hexagons/tuner"
	tunerports "hexagons/tuner/infrastructure"
)

// Implements tuner output ports towards hw, and hw output ports towards tuner
type TunerHwAdapter struct {
	tunerServicePort tunerports.TunerServicePort
	hwServicePort    hwports.HwServicePort
}

func NewTunerHwAdapter(tuner *tuner.TunerRoot, hw *hw.HwRoot) TunerHwAdapter {
	adapter := TunerHwAdapter{tuner.GetServicePort(), hw.GetServicePort()}
	return adapter
}

// Hw -> Tuner
func (adapter *TunerHwAdapter) UpdateStationList(stationList []string) {
	adapter.tunerServicePort.UpdateStationList(stationList)
}

// Hw -> Tuner
func (adapter *TunerHwAdapter) UpdateSubscription(subscription bool) {
	adapter.tunerServicePort.UpdateSubscription(subscription)
}

// Tuner -> Hw
func (adapter *TunerHwAdapter) TuneToStation(stationId uint32) {
	adapter.hwServicePort.TuneToStation(stationId)
}
