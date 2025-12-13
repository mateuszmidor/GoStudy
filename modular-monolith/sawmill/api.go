package sawmill

// Beam is a part of the sawmill module public API.
type Beam struct{}

// API of the sawmill module.
type API interface {
	GetBeams(count int) ([]Beam, error)
}
