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

func TestLRU_RecentGetPreventsEviction(t *testing.T) {
	// given
	lru := NewLRU(2)
	lru.Put("1", "a")
	lru.Put("2", "b")
	lru.Get("1") // access "1" to make it most recently used

	// when
	lru.Put("3", "c") // add "3", should evict "2"

	// then
	assert.Nil(t, lru.Get("2"), "expected '2' to be evicted")
	assert.NotNil(t, lru.Get("1"), "expected '1' to remain in cache")
	assert.NotNil(t, lru.Get("3"), "expected '3' to be present")
}

func TestPutSameKeyTwice_SizeIsOne(t *testing.T) {
	// given
	lru := NewLRU(2)

	// when
	lru.Put("key", "val1")
	lru.Put("key", "val2")

	// then
	assert.Equal(t, 1, lru.Size(), "Size() should be 1 after adding value under the same key twice")
	assert.Equal(t, 1, lru.tracking.Len(), "num of tracked elements should be 1")
	val := lru.Get("key")
	assert.NotNil(t, val, "Value should be present for 'key'")
	assert.Equal(t, "val2", *val, "Value should be updated to the most recent one")
}
