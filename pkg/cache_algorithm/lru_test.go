package cache_algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRU(t *testing.T) {
	var c LRUCache
	c.New(10)

	// 54321
	// edcba
	c.Set("1", []byte("a"))
	c.Set("2", []byte("b"))
	c.Set("3", []byte("c"))
	c.Set("4", []byte("d"))
	c.Set("5", []byte("e"))

	// 25431
	// bedca
	val, _ := c.Get("2")
	assert.EqualValues(t, []byte("b"), val)
	assert.EqualValues(t, "b", c.ll.Front().Value.(*lruEntry).value)

	// 62543
	// fbedc
	c.Set("6", []byte("f"))
	_, ok := c.Get("1")
	assert.EqualValues(t, false, ok)

	// 6254
	// fbed
	c.DelBack()
	assert.EqualValues(t, []byte("d"), c.ll.Back().Value.(*lruEntry).value)

	c.Set("2", []byte("a"))
	assert.EqualValues(t, []byte("a"), c.ll.Front().Value.(*lruEntry).value)

	c.Del("2")
	_, ok = c.cache["2"]
	assert.EqualValues(t, false, ok)
	assert.EqualValues(t, true, byte('b') != c.ll.Front().Value.(*lruEntry).value[0])
}
