package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestCreate_Success(t *testing.T) {
	// given
	db := NewMemDB()
	p := Product{Name: "shoes", Count: 2}

	// when
	created, err := db.Create(p)

	// then
	assert.NoError(t, err)
	assert.NotZero(t, created.ID)
}
