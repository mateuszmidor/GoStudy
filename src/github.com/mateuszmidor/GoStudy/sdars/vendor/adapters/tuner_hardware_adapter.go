package adapters

import (
	"actors/hardware"
	"hexagons/tuner"
	"hexagons/tuner/domain"
	"hexagons/tuner/application/cmds"
)

type HardwareAdapter struct {
	root *tuner.TunerRoot
	a *hardware.HwActor
}

func NewHardwareAdapter(r *tuner.TunerRoot, a *hardware.HwActor) HardwareAdapter {
	ha := HardwareAdapter{r, a}
	a.OnUpdatStationList = ha.UpdateStationList
	a.OnUpdateSubscription = ha.UpdateSubscription
	return ha
}

// Transform HW command into Tuner command
func (ha *HardwareAdapter) UpdateStationList(newStationList domain.StationList) {
	ha.root.CommandQueue <- cmds.NewUpdateStationListCmd(ha.root, newStationList)
}

// Transform HW command into Tuner command
func (ha *HardwareAdapter) UpdateSubscription(subscription domain.Subscription) {
	ha.root.CommandQueue <- cmds.NewUpdateSubscriptionCmd(ha.root, subscription)
}

// Transform Tuner command into HW command
func (ha HardwareAdapter) TuneToStation(stationId uint32) {
	// forward to HW actor
	ha.a.TuneToStation(stationId)
}
