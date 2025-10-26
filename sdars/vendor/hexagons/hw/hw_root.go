package hw

import "hexagons/hw/domain"
import "hexagons/hw/application"
import "hexagons/hw/infrastructure"

type HwRoot struct {
	state   domain.HwState
	service application.HwService
	ports   infrastructure.OuterWorldPorts
}

func NewHwRoot() HwRoot {
	root := HwRoot{}
	root.state = domain.NewHwState()
	root.service = application.NewHwService(&root.state)
	root.ports = infrastructure.OuterWorldPorts{}
	return root
}

func (root *HwRoot) SetTunerPort(port infrastructure.TunerPort) {
	root.ports.TunerPort = port
}

func (root *HwRoot) GetServicePort() infrastructure.HwServicePort {
	return &root.service
}

func (root *HwRoot) Run() {
	root.service.Run(&root.ports)
}
