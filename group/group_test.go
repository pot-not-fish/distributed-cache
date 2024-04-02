package group

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroup(t *testing.T) {
	var (
		g1 *Group
		g2 *Group
	)
	g1, err := NewGroup("cache_1", 6, func(key string) ([]byte, bool) {
		return g2.Get(key)
	}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	g2, err = NewGroup("cache_2", 6, func(key string) ([]byte, bool) {
		return g1.Get(key)
	}, nil)
	if err != nil {
		log.Fatalln(err)
	}

	g1.Set("a", []byte("1"))
	g2.Set("b", []byte("2"))
	val, ok := g1.Get("b")
	assert.EqualValues(t, true, ok)
	assert.EqualValues(t, []byte("2"), val)
}
