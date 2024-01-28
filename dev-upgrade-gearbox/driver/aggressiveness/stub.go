package aggressiveness

// Stub is unit test helper for those who need aggressiveness.Level
type Stub struct {
	Multiplier float64
}

// GetRPMMultiplier getter
func (s Stub) GetRPMMultiplier() float64 {
	return s.Multiplier
}
