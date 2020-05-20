package drivingmode

import (
	"driver/aggressiveness"
	"shared/gas"
	"shared/gear"
	"driver/types"
)

// Stub is unit test helper for those who need drivingmode.Level
type Stub struct {
	Min, Max types.RPM
	KickDown gear.Change
}

// GetOptimalRPM getter
func (s Stub) GetOptimalRPM(al aggressiveness.Level) (types.RPM, types.RPM) {
	return s.Min * al.GetRPMMultiplier(), s.Max * al.GetRPMMultiplier()
}

// GetKickDownGearChange getter
func (s Stub) GetKickDownGearChange(gas.Value) gear.Change {
	return s.KickDown
}
