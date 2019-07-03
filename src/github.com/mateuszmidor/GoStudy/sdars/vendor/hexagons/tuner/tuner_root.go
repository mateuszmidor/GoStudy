package tuner

import "hexagons/tuner/domain"
import "hexagons/tuner/application"
import "hexagons/tuner/infrastructure"

// Tuner aggregate root; visible to the outer world
type TunerRoot struct {
	tuner domain.Tuner
	ports infrastructure.Ports
	service application.TunerService
}

func NewTunerRoot() TunerRoot {
	root := TunerRoot{}
	root.tuner = domain.NewTuner()
	root.ports = infrastructure.Ports{}
	root.service = application.NewTunerService()
	return root
}

func (t *TunerRoot) SetupHwPortOut(hwPortOut infrastructure.HwPortOut) {
	t.ports.HwPortOut = hwPortOut
}

func (t *TunerRoot) SetupUiPortOut(uiPortOut infrastructure.UiPortOut) {
	t.ports.UiPortOut = uiPortOut
}

func (t *TunerRoot) GetUiPortIn() infrastructure.UiPortIn {
	return &t.service // TunerService implements all the input ports
}

func (t *TunerRoot) GetHwPortIn() infrastructure.HwPortIn {
	return &t.service // TunerService implements all the input ports
}

// To be run from non-main gorutine
func (t *TunerRoot) Run() {
	t.service.Run(&t.tuner, &t.ports)
}