package externalsystemsfacade

import (
	"externalsystems"
	"driver/types"
)

// Facade hides ExernalSystems atrocity
type Facade struct {
	ExternalSystems   *externalsystems.ExternalSystems
	IsTrailorAttached bool
}

// GetDrivingDownTheSlope tells if the car is going down the slope
func (f Facade) GetDrivingDownTheSlope() bool {
	ligts := f.ExternalSystems.GetLights()
	position := ligts.GetPosition()
	return position != nil && *position >= 7 && *position <= 10
}

// GetAngularSpeed is getter
func (f Facade) GetAngularSpeed() types.AngularSpeed {
	return f.ExternalSystems.GetAngularSpeed()
}

// GetCurrentRPM is getter
func (f Facade) GetCurrentRPM() types.RPM {
	return f.ExternalSystems.GetCurrentRpm()
}

// GetTrailorAttached is getter
func (f Facade) GetTrailorAttached() bool {
	return f.IsTrailorAttached
}

// IsManualGearChangeActive getter
func (f Facade) IsManualGearChangeActive() bool {
	return false // TODO: should get this info from somewhere
}
