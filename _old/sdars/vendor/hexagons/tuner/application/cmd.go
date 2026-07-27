package application

import "hexagons/tuner/domain"
import "hexagons/tuner/infrastructure"

type Cmd interface {
	Execute(state *domain.TunerState, ports *infrastructure.OuterWorldPorts)
}
