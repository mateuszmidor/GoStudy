package main

import (
	"hexagons/tuner"
	tunerports "hexagons/tuner/infrastructure"
	"hexagons/ui"
	uiports "hexagons/ui/infrastructure"
)

// TunerUIAdapter implements tuner output ports towards ui, and ui output ports towards tuner
type TunerUIAdapter struct {
	tunerServicePort tunerports.TunerServicePort
	uiServicePort    uiports.ServicePort
}

// NewTunerUIAdapter creates a Tuner - Ui bridge
func NewTunerUIAdapter(tuner *tuner.TunerRoot, ui *ui.UiRoot) TunerUIAdapter {
	adapter := TunerUIAdapter{tuner.GetServicePort(), ui.GetServicePort()}
	return adapter
}

// TuneToStation forwards command UI -> Tuner
func (adapter *TunerUIAdapter) TuneToStation(stationId uint32) {
	adapter.tunerServicePort.TuneToStation(stationId)
}

// UpdateStationList forwards command Tuner -> UI
func (adapter *TunerUIAdapter) UpdateStationList(stationList []string) {
	adapter.uiServicePort.UpdateStationList(stationList)
}

// UpdateSubscription forwards command Tuner -> UI
func (adapter *TunerUIAdapter) UpdateSubscription(active bool) {
	adapter.uiServicePort.UpdateSubscription(active)
}
