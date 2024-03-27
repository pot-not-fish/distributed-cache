package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type str string

func (s str) Len() int {
	return len(s)
}

func TestSum(t *testing.T) {
	c := New(10, nil)

	// 54321
	// edcba
	c.Set("1", str("a"))
	c.Set("2", str("b"))
	c.Set("3", str("c"))
	c.Set("4", str("d"))
	c.Set("5", str("e"))

	// 25431
	// bedca
	val, _ := c.Get("2")
	assert.EqualValues(t, "b", val)
	assert.EqualValues(t, "b", c.ll.Front().Value.(*entry).value)

	// 62543
	// fbedc
	c.Set("6", str("f"))
	_, ok := c.Get("1")
	assert.EqualValues(t, false, ok)

	// 6254
	// fbed
	c.Del()
	assert.EqualValues(t, "d", c.ll.Back().Value.(*entry).value)
}
