package gearadjustment

import (
	"shared/gear"
	"driver/types"
)

// AdjustForOptimalRPM calculates optimal gear change for current RPM
func AdjustForOptimalRPM(current, min, max types.RPM) gear.Change {
	switch {
	case current < min:
		return gear.GearDown
	case current > max:
		return gear.GearUp
	default:
		return gear.KeepCurrent
	}
}
