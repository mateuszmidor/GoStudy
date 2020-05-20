package aggressiveness

// Level represents gearbox aggressiveness level
type Level interface {
	GetRPMMultiplier() float64
}
