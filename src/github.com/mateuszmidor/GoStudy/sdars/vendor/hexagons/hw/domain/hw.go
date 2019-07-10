package domain

type Hw struct {
	CurrentStationId uint32
}

func NewHw() Hw {
	return Hw{}
}