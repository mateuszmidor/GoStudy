package util

import "math"

const epsilon = 0.001

// IsEqual compares floats with delta
func IsEqual(l, r float64) bool {
	return math.Abs(l-r) < epsilon
}
