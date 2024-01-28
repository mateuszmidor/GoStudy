package aggressiveness

// Level1 represents gearbox aggressiveness level 1
type Level1 struct {
}

// NewLevel1 is constructor
func NewLevel1() Level1 {
	return Level1{}
}

// GetRPMMultiplier getter
func (l Level1) GetRPMMultiplier() float64 {
	return 1.0
}
