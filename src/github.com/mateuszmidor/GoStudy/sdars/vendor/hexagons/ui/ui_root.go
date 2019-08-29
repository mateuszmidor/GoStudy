package ui

import "hexagons/ui/domain"
import "hexagons/ui/application"
import "hexagons/ui/infrastructure"

type UiRoot struct {
	state   domain.UiState
	service application.UiService
	ports   infrastructure.OuterWorldPorts
}

func NewUiRoot() UiRoot {
	root := UiRoot{}
	root.state = domain.NewUiState()
	root.service = application.NewUiService(&root.state)
	root.ports = infrastructure.OuterWorldPorts{}
	return root
}

func (root *UiRoot) SetTunerPort(tuner infrastructure.TunerPort) {
	root.ports.TunerPort = tuner
}

func (root *UiRoot) GetServicePort() infrastructure.UiServicePort {
	return &root.service
}

func (root *UiRoot) Run() {
	root.service.Run(&root.ports)
}
