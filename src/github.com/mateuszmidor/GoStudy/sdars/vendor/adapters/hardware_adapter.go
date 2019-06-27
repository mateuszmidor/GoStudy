package adapters

import (
	"sdars"
	"actors/hardware"
	"cmds"
)

type HardwareAdapter struct {
	commandQueue *sdars.CommandQueue
	a *hardware.HwActor
}

func NewHardwareAdapter(c *sdars.CommandQueue, a *hardware.HwActor) HardwareAdapter {
	ha := HardwareAdapter{c, a}
	a.OnUpdatStationList = ha.UpdateStationList
	a.OnUpdateSubscription = ha.UpdateSubscription
	return ha
}

// Transform HW command into Tuner command
func (ha *HardwareAdapter) UpdateStationList(newStationList []string) {
	*ha.commandQueue <- cmds.NewUpdateStationListCmd(newStationList)
}

// Transform HW command into Tuner command
func (ha *HardwareAdapter) UpdateSubscription(active bool) {
	*ha.commandQueue <- cmds.NewUpdateSubscriptionCmd(active)
}

// Transform Tuner command into HW command
func (ha HardwareAdapter) TuneToStation(stationId uint32) {
	// forward to HW actor
	ha.a.TuneToStation(stationId)
}
