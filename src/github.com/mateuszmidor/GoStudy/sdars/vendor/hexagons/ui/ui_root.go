package ui

import "hexagons/ui/domain"
import "hexagons/ui/application"
import "hexagons/ui/infrastructure"

type UiRoot struct {
	ui domain.Ui
	service application.UiService
	ports infrastructure.Ports
}

func NewUiRoot() UiRoot {
	root := UiRoot{}
	root.ui = domain.NewUi()
	root.service = application.NewUiService(&root.ui)
	root.ports = infrastructure.Ports{}
	return root
}

func (root *UiRoot) SetupTunerPortOut(tuner infrastructure.TunerPortOut) {
	root.ports.TunerPortOut = tuner
}

func (root *UiRoot) GetTunerPortIn() infrastructure.TunerPortIn {
	return &root.service // UiService implements all the input ports 
}

func (root *UiRoot) Run() {
	root.service.Run(root.ports)
}