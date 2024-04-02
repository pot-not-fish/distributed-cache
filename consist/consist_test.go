package consist

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsist(t *testing.T) {
	m := New(2, func(data []byte) uint32 {
		a1, _ := strconv.ParseInt(string(data[0]), 10, 64)
		a2, _ := strconv.ParseInt(string(data[1:]), 10, 64)
		return uint32(100*a2 + a1*100)
	})

	m.Set("10")
	m.Set("20")
	m.Set("30")

	assert.EqualValues(t, m.nodes, []int{1000, 1100, 1200, 2000, 2100, 2200, 3000, 3100, 3200})

	assert.EqualValues(t, m.Get("111"), "10")
	assert.EqualValues(t, m.Get("221"), "30")
	assert.EqualValues(t, m.Get("150"), "10")
}
