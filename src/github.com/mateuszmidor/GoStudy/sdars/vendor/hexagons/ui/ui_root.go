package ui

import "hexagons/ui/domain"
import "hexagons/ui/application"
import "hexagons/ui/infrastructure"

type UiRoot struct {
	ui      domain.Ui
	service application.UiService
	ports   infrastructure.Ports
}

func NewUiRoot() UiRoot {
	root := UiRoot{}
	root.ui = domain.NewUi()
	root.service = application.NewUiService(&root.ui)
	root.ports = infrastructure.Ports{}
	return root
}

func (root *UiRoot) SetTunerPort(tuner infrastructure.TunerCommandsPort) {
	root.ports.CommandsPort = tuner
}

func (root *UiRoot) GetServicePort() infrastructure.ServicePort {
	return &root.service
}

func (root *UiRoot) Run() {
	root.service.Run(&root.ports)
}
