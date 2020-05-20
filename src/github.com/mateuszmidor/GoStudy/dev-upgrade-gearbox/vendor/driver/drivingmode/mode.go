package drivingmode

import (
	"driver/aggressiveness"
	"shared/gas"
	"shared/gear"
	"driver/types"
)

// Mode represents a driving mode like eco/comfort/sport
type Mode interface {
	GetOptimalRPM(aggressiveness.Level) (types.RPM, types.RPM)
	GetKickDownGearChange(gas gas.Value) gear.Change
}
