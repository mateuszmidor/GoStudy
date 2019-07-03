package adapters

import (
	"hexagons/ui"
	uiports "hexagons/ui/infrastructure"
	"hexagons/tuner"
	tunerports "hexagons/tuner/infrastructure"
)

type UiAdapter struct {
	tunerPortIn tunerports.UiPortIn
	uiPortIn uiports.TunerPortIn
}

func NewUiAdapter(tuner *tuner.TunerRoot, ui *ui.UiRoot) UiAdapter {
	ua := UiAdapter{tuner.GetUiPortIn(), ui.GetTunerPortIn()}
	return ua
}

// UI -> Tuner
func (adapter UiAdapter) TuneToStation(stationId uint32) {
	adapter.tunerPortIn.TuneToStation(stationId)
}

// Tuner -> UI
func (adapter UiAdapter) UpdateStationList(stationList []string) {
	// forward to cluster actor
	adapter.uiPortIn.UpdateStationList(stationList)
}

// Tuner -> UI
func (adapter UiAdapter) UpdateSubscription(active bool) {
	// forward to cluster actor
	adapter.uiPortIn.UpdateSubscription(active)	
}