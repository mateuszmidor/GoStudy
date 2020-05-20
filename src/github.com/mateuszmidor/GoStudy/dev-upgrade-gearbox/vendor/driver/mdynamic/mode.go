package mdynamic

import (
	"driver/uślizg"
	"driver/types"
)

// Mode represents MDynamic mode
type Mode struct {
	enabled bool
}

// Enable enables MDynamic mode
func (m *Mode) Enable() {
	m.enabled = true
}

// IsGearChangeAllowed does what it says
func (m *Mode) IsGearChangeAllowed(as types.AngularSpeed) bool {
	return !m.enabled || !uślizg.IsUślizg(as)
}
