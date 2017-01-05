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
        //expand to 100
        for i := 0; i < 80; i++ {
                item := fmt.Sprintf("hello%v", i)
                q.Push(item)
        }
        //head is now at 40
        for i := 0; i < 40; i++ {
                item := fmt.Sprintf("hello%v", i)
                res, _ := q.Pop()
                assert.Equal(t, item, res)
        }
        //make sure tail wraps around => tail is at (80+50)%100=30
        for i := 0; i < 50; i++ {
                item := fmt.Sprintf("hello%v", i+80)
                q.Push(item)
        }
        //now pop enough to make the head wrap around => (40 + 80)%100=20
        for i := 0; i < 80; i++ {
                item := fmt.Sprintf("hello%v", i+40)
                res, _ := q.Pop()
                assert.Equal(t, item, res)
        }
        //push enough to cause expansion
        for i := 0; i < 100; i++ {
                item := fmt.Sprintf("hello%v", i+130)
                q.Push(item)
        }
        //empty the queue
        for i := 0; i < 110; i++ {
                item := fmt.Sprintf("hello%v", i+120)
                res, _ := q.Pop()
                assert.Equal(t, item, res)
        }
        assert.True(t, q.Empty())
}

