package adapters

import (
	"hexagons/hw"
	"hexagons/tuner"
	hwports "hexagons/hw/infrastructure"
	tunerports "hexagons/tuner/infrastructure"
)

type HwAdapter struct {
	tunerInPort tunerports.HwPortIn
	hwInPort hwports.TunerInPort
}

func NewHwAdapter(tuner *tuner.TunerRoot, hw *hw.HwRoot) HwAdapter {
	adapter := HwAdapter{tuner.GetHwPortIn(), hw.GetTunerInPort()}
	return adapter
}

// Hw -> Tuner
func (adapter *HwAdapter) UpdateStationList(stationList []string) {
	adapter.tunerInPort.StationListUpdated(stationList) 
}

// Hw -> Tuner
func (adapter *HwAdapter) UpdateSubscription(subscription bool) {
	adapter.tunerInPort.SubscriptionUpdated(subscription)
}

// Tuner -> Hw
func (adapter *HwAdapter) TuneToStation(stationId uint32) {
	adapter.hwInPort.TuneToStation(stationId)
}
