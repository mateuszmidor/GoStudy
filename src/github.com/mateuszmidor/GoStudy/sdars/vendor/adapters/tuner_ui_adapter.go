package adapters

import (
	"actors/ui"
	"hexagons/tuner"
	"hexagons/tuner/domain"
	"hexagons/tuner/application/cmds"
)

type UiAdapter struct {
	root *tuner.TunerRoot
	a *ui.UiActor
}

func NewUiAdapter(r *tuner.TunerRoot, a *ui.UiActor) UiAdapter {
	ua := UiAdapter{r, a}
	a.OnTuneToStation = ua.TuneToStation
	return ua
}

// Transform Cluster command into Tuner command
func (ua *UiAdapter) TuneToStation(stationId domain.StationId) {
	ua.root.PutCommand(cmds.NewTuneToStationCmd(stationId))
}

// Transform Tuner command into Cluster command
func (ua UiAdapter) UpdateStationList(stationList []string) {
	// forward to cluster actor
	ua.a.UpdateStationList(stationList)
}

// Transform Tuner command into Cluster command
func (ua UiAdapter) UpdateSubscription(active bool) {
	// forward to cluster actor
	ua.a.UpdateSubscription(active)	
}