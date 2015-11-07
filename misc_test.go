package misc

import (
	"github.com/stvp/assert"
	"testing"
)

func Test_CompareStrings(t *testing.T) {
	assert.Equal(t, 1, CompareStrings("hi", "bye"), "compare 1 failed")
	assert.Equal(t, -1, CompareStrings("a", "b"), "compare 2 failed")
	assert.Equal(t, 0, CompareStrings("x", "x"), "compare 3 failed")
}
