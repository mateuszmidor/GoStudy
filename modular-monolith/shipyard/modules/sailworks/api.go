package sailworks

// Sail is a part of the ropeworks public API.
type Sail struct{}

// API of the sailworks module.
type API interface {
	GetSails(count int) ([]Sail, error)
}
