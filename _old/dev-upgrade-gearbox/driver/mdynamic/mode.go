package mdynamic

import (
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/driver/types"
	"github.com/mateuszmidor/GoStudy/dev-upgrade-gearbox/driver/uslizg"
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
	return !m.enabled || !uslizg.IsUślizg(as)
}
