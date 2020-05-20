package gas

import (
	"fmt"
)

// Value represents accelerator pedal state 0..1
type Value struct {
	value float64
}

// Zero represents totally release gas pedal
var Zero = New(0.0)

// Half represents half-pressed gas pedal
var Half = New(0.5)

// Full represents fully pressed gas pedal
var Full = New(1.0)

// New creates new Gas Value
func New(value float64) Value {
	if value < 0.0 || value > 1.0 {
		panic(fmt.Sprintf("Gas value must be in 0..1 range, was %f", value))
	}
	return Value{value: value}
}

// ReachedThreshold checks if v >= t
func (v Value) ReachedThreshold(t Threshold) bool {
	return v.value >= t.value
}

// IsZero says if gas is released
func (v Value) IsZero() bool {
	return v.value < 0.001
}

// IsFull says if gas is pushed to the limit
func (v Value) IsFull() bool {
	return v.value > 0.999
}
