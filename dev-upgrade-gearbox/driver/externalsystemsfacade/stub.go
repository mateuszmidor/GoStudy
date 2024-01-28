package externalsystemsfacade

import (
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/driver/types"
)

// Stub is unit test helper for those who need externalsystemsfacade.Data
type Stub struct {
	DrivingDownTheSlope    bool
	TrailorAttached        bool
	AngularSpeed           types.AngularSpeed
	CurrentRPM             types.RPM
	ManualGearChangeActive bool
}

// SetDownTheSlope setter
func (s Stub) SetDownTheSlope() Stub {
	s.DrivingDownTheSlope = true
	return s
}

// SetTrailor setter
func (s Stub) SetTrailor() Stub {
	s.TrailorAttached = true
	return s
}

// SetUslizg setter
func (s Stub) SetUslizg(angularSpeedForUślizg float64) Stub {
	s.AngularSpeed = angularSpeedForUślizg
	return s
}

// SetRPM setter
func (s Stub) SetRPM(rpm types.RPM) Stub {
	s.CurrentRPM = rpm
	return s
}

// SetManualGearChangeActive setter
func (s Stub) SetManualGearChangeActive() Stub {
	s.ManualGearChangeActive = true
	return s
}

// IsManualGearChangeActive getter
func (s Stub) IsManualGearChangeActive() bool {
	return s.ManualGearChangeActive
}

// GetDrivingDownTheSlope tells if the car is going down the slope
func (s Stub) GetDrivingDownTheSlope() bool {
	return s.DrivingDownTheSlope
}

// GetAngularSpeed is getter
func (s Stub) GetAngularSpeed() types.AngularSpeed {
	return s.AngularSpeed
}

// GetCurrentRPM is getter
func (s Stub) GetCurrentRPM() types.RPM {
	return s.CurrentRPM
}

// GetTrailorAttached is getter
func (s Stub) GetTrailorAttached() bool {
	return s.TrailorAttached
}
