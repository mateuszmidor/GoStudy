package externalsystems

// ExternalSystems represents data from outside the GearBoxDriver
type ExternalSystems struct {
	currentRpm   float64
	angularSpeed float64
	lights       Lights
}

// NewExternalSystems is constructor
func NewExternalSystems(currentRpm, angularSpeed float64, lights Lights) ExternalSystems {
	return ExternalSystems{currentRpm, angularSpeed, lights}
}

// GetAngularSpeed getter
func (es ExternalSystems) GetAngularSpeed() float64 {
	return es.angularSpeed
}

// GetCurrentRpm getter
func (es ExternalSystems) GetCurrentRpm() float64 {
	return es.currentRpm
}

// GetLights getter
func (es ExternalSystems) GetLights() Lights {
	return es.lights
}
