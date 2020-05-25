package aggressiveness

// Level3 represents gearbox aggressiveness level 3
type Level3 struct {
}

// NewLevel3 is constructor
func NewLevel3() Level3 {
	return Level3{}
}

// GetRPMMultiplier getter
func (l Level3) GetRPMMultiplier() float64 {
	return 1.3
}
