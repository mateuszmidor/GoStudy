package tuner

import "hexagons/tuner/domain"
import "hexagons/tuner/application"
import "hexagons/tuner/infrastructure"

// TunerRoot; aggregate visible to the outer world
type TunerRoot struct {
	state   domain.TunerState              // keeps internal state
	ports   infrastructure.OuterWorldPorts // allows tuner talking to outer world
	service application.TunerService       // allows outer world talking to tuner
}

func NewTunerRoot() TunerRoot {
	return TunerRoot{domain.NewTunerState(), infrastructure.OuterWorldPorts{}, application.NewTunerService()}
}

func (t *TunerRoot) SetHwPort(hwPortOut infrastructure.HwPort) {
	t.ports.HwPort = hwPortOut
}

func (t *TunerRoot) SetUiPort(uiPortOut infrastructure.UIPort) {
	t.ports.UIPort = uiPortOut
}

func (t *TunerRoot) GetServicePort() infrastructure.TunerServicePort {
	return &t.service
}

// To be run from non-main gorutine
func (t *TunerRoot) Run() {
	t.service.Run(&t.state, &t.ports)
}
