package drivingmode

import (
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/driver/aggressiveness"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/driver/types"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/shared/gas"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/shared/gear"
)

// Mode represents a driving mode like eco/comfort/sport
type Mode interface {
	GetOptimalRPM(aggressiveness.Level) (types.RPM, types.RPM)
	GetKickDownGearChange(gas gas.Value) gear.Change
}
