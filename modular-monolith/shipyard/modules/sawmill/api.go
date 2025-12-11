package sawmill

// Plank is a part of the sawmill module public API.
type Plank struct{}

// API of the sawmill module.
type API interface {
	Run()
	GetPlanks(count int) ([]Plank, error)
}
