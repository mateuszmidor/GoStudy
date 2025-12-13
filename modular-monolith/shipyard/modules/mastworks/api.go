package mastworks

// Mast is a part of the mastworks module public API.
type Mast struct{}

// API of the mastworks module.
type API interface {
	GetMasts(count int) ([]Mast, error)
}
