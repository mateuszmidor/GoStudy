package gear

import "fmt"

// Change represents gear change value
type Change int

// GearUp says: increase current gear
const GearUp Change = 1

// KeepCurrent says: keep current gear
const KeepCurrent Change = 0

// GearDown says: decrease current gear
const GearDown Change = -1

// DoubleGearDown says: decrease current gear
const DoubleGearDown Change = -2

// Add returns a sum of two gear changes eg -1 + -1 -> -2
func (gc Change) Add(other Change) Change {
	return gc + other
}
func (gc Change) String() string {
	switch {
	case gc < 0:
		return fmt.Sprintf("Gear down by %d", -gc)
	case gc > 0:
		return fmt.Sprintf("Gear up by %d", gc)
	default:
		return "Keep current gear"
	}
}
