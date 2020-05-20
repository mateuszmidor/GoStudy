package externalsystemsfacade

import "driver/types"

// Data is interface for getting external systems data
type Data interface {
	GetDrivingDownTheSlope() bool
	GetTrailorAttached() bool
	GetAngularSpeed() types.AngularSpeed
	GetCurrentRPM() types.RPM
	IsManualGearChangeActive() bool
}
