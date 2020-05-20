package drivingmode

import (
	"driver/aggressiveness"
	"shared/gear"
	"driver/types"
	"shared/gas"
)

// Eco represents economic driving mode
type Eco struct {
	min, max types.RPM
}

// NewEco is constructor
func NewEco(min, max types.RPM) Eco {
	return Eco{min, max}
}

// GetOptimalRPM getter
func (e Eco) GetOptimalRPM(aggressiveness.Level) (types.RPM, types.RPM) {
	return e.min, e.max
}

// GetKickDownGearChange getter
func (e Eco) GetKickDownGearChange(gas gas.Value) gear.Change {
	return gear.KeepCurrent // eco mode doesnt allow for kickdown
}
