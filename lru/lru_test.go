package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	var c Cache
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
	assert.EqualValues(t, "b", c.ll.Front().Value.(*entry).value)

	// 62543
	// fbedc
	c.Set("6", []byte("f"))
	_, ok := c.Get("1")
	assert.EqualValues(t, false, ok)

	// 6254
	// fbed
	c.Del()
	assert.EqualValues(t, []byte("d"), c.ll.Back().Value.(*entry).value)

	c.Set("2", []byte("a"))
	assert.EqualValues(t, []byte("a"), c.ll.Front().Value.(*entry).value)
}
