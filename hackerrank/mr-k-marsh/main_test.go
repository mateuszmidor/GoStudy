package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getMaxPerimeter_possible(t *testing.T) {
	// given
	grid := []string{"...", ".x.", "..."}

	// when
	perimeter := getMaxPerimeter(grid, 3, 3)

	// then
	assert.Equal(t, 8, perimeter)
}

func Test_getMaxPerimeter_2_fields_possible(t *testing.T) {
	// given
	grid := []string{"..xx",
		"..xx",
		"....",
		".x..",
		"...."}

	// when
	perimeter := getMaxPerimeter(grid, 4, 5)

	// then
	assert.Equal(t, 10, perimeter)
}

func Test_getMaxPerimeter_impossible(t *testing.T) {
	// given
	grid := []string{"x..", ".x.", "..x"}

	// when
	perimeter := getMaxPerimeter(grid, 3, 3)

	// then
	assert.Equal(t, -1, perimeter)
}
