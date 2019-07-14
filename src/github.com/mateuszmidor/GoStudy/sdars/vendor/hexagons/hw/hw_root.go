package hw

import "hexagons/hw/domain"
import "hexagons/hw/application"
import "hexagons/hw/infrastructure"

type HwRoot struct {
	hw      domain.Hw
	service application.HwService
	ports   infrastructure.Ports
}

func NewHwRoot() HwRoot {
	root := HwRoot{}
	root.hw = domain.NewHw()
	root.service = application.NewHwService(&root.hw)
	root.ports = infrastructure.Ports{}
	return root
}

func (root *HwRoot) SetTunerPort(port infrastructure.EventsPort) {
	root.ports.EventsPort = port
}

func (root *HwRoot) GetServicePort() infrastructure.ServicePort {
	return &root.service
}

func (root *HwRoot) Run() {
	root.service.Run(&root.ports)
}
