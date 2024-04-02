package consist

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsist(t *testing.T) {
	m := New(2, func(data []byte) uint32 {
		a1, _ := strconv.ParseInt(string(data[0]), 10, 64)
		a2, _ := strconv.ParseInt(string(data[1]), 10, 64)
		return uint32(100*a2 + a1*10)
	})

	m.Set("1")
	m.Set("2")
	m.Set("3")

	assert.EqualValues(t, m.nodes, []int{100, 110, 120, 200, 210, 220, 300, 310, 320})
}
