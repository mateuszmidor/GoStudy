package externalsystemsfacade

import "github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/driver/types"

// Data is interface for getting external systems data
type Data interface {
	GetDrivingDownTheSlope() bool
	GetTrailorAttached() bool
	GetAngularSpeed() types.AngularSpeed
	GetCurrentRPM() types.RPM
	IsManualGearChangeActive() bool
}
