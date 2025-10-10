package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSimplePutGet(t *testing.T) {
	// given
	lru := NewLRU(2)

	// when
	lru.Put("1", "a")

	// then
	val := lru.Get("1")
	assert.NotNil(t, val)
	assert.Equal(t, "a", *val)
}

func TestExceedCapacityDiscardsOldertItem(t *testing.T) {
	// given
	lru := NewLRU(2)

	// when
	lru.Put("1", "a") // this one should get discarded
	time.Sleep(time.Millisecond * 500)
	lru.Put("2", "b")
	time.Sleep(time.Millisecond * 500)
	lru.Put("3", "c")

	// then
	assert.Nil(t, lru.Get("1"))
	assert.NotNil(t, lru.Get("2"))
	assert.NotNil(t, lru.Get("3"))
}
