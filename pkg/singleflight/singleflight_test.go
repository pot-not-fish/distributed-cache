package singleflight

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleflight(t *testing.T) {
	var a = map[string][]byte{
		"1": []byte("a"),
		"2": []byte("b"),
		"3": []byte("c"),
	}

	g := new(Group)

	var (
		b    []byte
		err  error
		cnt1 = 0
	)
	for i := 0; i < 10; i++ {
		go func() {
			b, err = g.Do("2", func(key string) ([]byte, error) {
				cnt1++
				if val, ok := a[key]; ok {
					return []byte(val), nil
				} else {
					return nil, fmt.Errorf("nil value")
				}
			})
			assert.EqualValues(t, b, []byte("b"))
			assert.EqualValues(t, err, nil)
		}()
	}

	cnt2 := 0
	for i := 0; i < 10; i++ {
		b, err = g.Do("2", func(key string) ([]byte, error) {
			cnt2++
			if val, ok := a[key]; ok {
				return []byte(val), nil
			} else {
				return nil, fmt.Errorf("nil value")
			}
		})
		assert.EqualValues(t, b, []byte("b"))
		assert.EqualValues(t, err, nil)
	}
	assert.EqualValues(t, b, []byte("b"))
	assert.EqualValues(t, err, nil)
	assert.EqualValues(t, cnt2 > cnt1, true)
}
