package testify_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssert(t *testing.T) {
	// "assert" reports error and continues testing function
	assert := assert.New(t)

	assert.Nil(nil)
	assert.NotNil(nil)

	assert.Equal(69, 96)
	assert.NotEqual(69, 96)

	assert.Contains(t, "abc", "d")                               // string contains substring
	assert.Contains([]int{1, 2, 3}, 5)                           // slice contains element
	assert.Contains(map[string]int{"aga": 22, "jan": 33}, "ola") // map contains key

	assert.ElementsMatch([]int{1, 2, 3}, []int{3, 2}) // collections contain same elements; ignore order

	assert.Same(new(int), new(int)) // pointers reference same object in memory

	assert.FailNow("Reached end of testing") // fatal error; terminate testing function
}

func TestRequire(t *testing.T) {
	// "require" reports error and terminates testing function; "fatal error"
	require := require.New(t)

	require.Nil(nil)
	require.NotNil(nil)

	require.FailNow("Reached end of testing")
}
