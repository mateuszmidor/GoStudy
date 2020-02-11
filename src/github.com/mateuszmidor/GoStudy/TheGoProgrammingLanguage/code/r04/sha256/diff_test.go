package main

import "testing"

type data struct {
	a, b         byte
	numDifferent int
}

const calcDiffBitsFailureString = "For numbers %d and %d the expected num of diff bits was %d but calculated %d"

func TestCalcDiffBits(t *testing.T) {
	in := []data{
		{0, 1, 1},
		{0, 2, 1},
		{0, 3, 2},
		{0, 255, 8},
	}

	for _, tc := range in {
		if actualDifferent := calcDiffBits(tc.a, tc.b); actualDifferent != tc.numDifferent {
			t.Errorf(calcDiffBitsFailureString, tc.a, tc.b, tc.numDifferent, actualDifferent)
		}
	}
}
