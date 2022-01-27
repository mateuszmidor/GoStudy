//go:generate stringer -type=Pill
package main

type Pill int

const (
	Ibuprom Pill = iota
	Paracetamol
	Aspirin
	C_Vitamin
)
