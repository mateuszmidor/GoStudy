package ropeworks

// Rope is a part of the ropeworks public API.
type Rope struct{}

// API of the ropeworks module.
type API interface {
	GetRopes(count int) ([]Rope, error)
}
