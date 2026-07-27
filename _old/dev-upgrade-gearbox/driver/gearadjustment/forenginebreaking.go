package gearadjustment

import "github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/shared/gear"

// AdjustGearForEngineBreaking calculates optimal gear change for engine breaking
func AdjustGearForEngineBreaking(isDownSlope, isTrailorAttached bool) gear.Change {
	if !isDownSlope {
		return gear.KeepCurrent
	}

	if !isTrailorAttached {
		return gear.KeepCurrent
	}

	return gear.GearDown
}
