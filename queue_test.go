package goring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushPop(t *testing.T) {
	q := New(10)
	q.Push("hello")
	res, _ := q.Pop()
	assert.Equal(t, "hello", res)
	assert.True(t, q.Empty())
}

func TestPushPopRepeated(t *testing.T) {
	q := New(10)
	for i := 0; i < 100; i++ {
		q.Push("hello")
		res, _ := q.Pop()
		assert.Equal(t, "hello", res)
		assert.True(t, q.Empty())
	}
}

func TestPushPopMany(t *testing.T) {
	q := New(10)
	for i := 0; i < 10000; i++ {
		item := fmt.Sprintf("hello%v", i)
		q.Push(item)
		res, _ := q.Pop()
		assert.Equal(t, item, res)
	}
	assert.True(t, q.Empty())
}

func TestPushPopMany2(t *testing.T) {
	q := New(10)
	for i := 0; i < 10000; i++ {
		item := fmt.Sprintf("hello%v", i)
		q.Push(item)
	}
	for i := 0; i < 10000; i++ {
		item := fmt.Sprintf("hello%v", i)
		res, _ := q.Pop()
		assert.Equal(t, item, res)
	}
	assert.True(t, q.Empty())
}
func TestExpand(t *testing.T) {
        q := New(10)
        for i := 0; i < 90000; i++ {
                item := fmt.Sprintf("hello%v", i)
                q.Push(item)
        }
        for i := 0; i < 1000; i++ {
                item := fmt.Sprintf("hello%v", i)
                res, _ := q.Pop()
                assert.Equal(t, item, res)
        }
        for i := 0; i < 500; i++ {
                item := fmt.Sprintf("hello%v", i+9000)
                q.Push(item)
        }
        for i := 0; i < 8500; i++ {
                item := fmt.Sprintf("hello%v", i+1000)
                res, _ := q.Pop()
                assert.Equal(t, item, res)
        }

        assert.True(t, q.Empty())
}

