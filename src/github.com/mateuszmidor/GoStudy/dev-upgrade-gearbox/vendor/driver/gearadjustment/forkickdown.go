package gearadjustment

import (
	"driver/drivingmode"
	"shared/gas"
	"shared/gear"
)

// AdjustForKickDown calculates optimal gear change for current given DriveMode
func AdjustForKickDown(v drivingmode.Mode, gas gas.Value) gear.Change {
	return v.GetKickDownGearChange(gas)
}
