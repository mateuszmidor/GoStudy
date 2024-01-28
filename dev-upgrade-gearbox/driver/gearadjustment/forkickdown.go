package gearadjustment

import (
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/driver/drivingmode"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/shared/gas"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/shared/gear"
)

// AdjustForKickDown calculates optimal gear change for current given DriveMode
func AdjustForKickDown(v drivingmode.Mode, gas gas.Value) gear.Change {
	return v.GetKickDownGearChange(gas)
}
