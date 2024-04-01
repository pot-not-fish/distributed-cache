package mutex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMutex(t *testing.T) {
	cache := New(6)

	// 3 2 1
	// c b a
	cache.Set("1", []byte("a"))
	cache.Set("2", []byte("b"))
	cache.Set("3", []byte("c"))

	cache.Get("1")
	cache.Set("4", []byte("d"))
	_, ok := cache.Get("2")
	assert.EqualValues(t, false, ok)
}
