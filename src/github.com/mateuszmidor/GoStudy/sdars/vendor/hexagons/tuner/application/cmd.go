package application

import "hexagons/tuner/domain"
import "hexagons/tuner/infrastructure"

type Cmd interface {
	Execute(tuner *domain.Tuner, ports *infrastructure.Ports)
}