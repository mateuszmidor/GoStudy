package drivingmode

import (
	"driver/aggressiveness"
	"shared/gas"
	"shared/gear"
	"driver/types"
)

// Sport represents sport driving mode
type Sport struct {
	min, max             types.RPM
	kickDown1, kickDown2 gas.Threshold
}

// NewSport is constructor
func NewSport(min, max types.RPM, kickdown1, kickdown2 gas.Threshold) Sport {
	return Sport{min, max, kickdown1, kickdown2}
}

// GetOptimalRPM getter
func (e Sport) GetOptimalRPM(a aggressiveness.Level) (types.RPM, types.RPM) {
	return e.min * a.GetRPMMultiplier(), e.max * a.GetRPMMultiplier()
}

// GetKickDownGearChange getter
func (e Sport) GetKickDownGearChange(gas gas.Value) gear.Change {
	switch {
	case gas.ReachedThreshold(e.kickDown2):
		return gear.DoubleGearDown
	case gas.ReachedThreshold(e.kickDown1):
		return gear.GearDown
	default:
		return gear.KeepCurrent
	}
}
