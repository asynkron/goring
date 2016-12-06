package goring

import (
	"sync"
	"sync/atomic"
)

type ringBuffer struct {
	buffer []interface{}
	head   uint64
	tail   uint64
	mod    uint64
}

type Queue struct {
	content *ringBuffer
	len     uint64
	lock    sync.Mutex
}

func New(initialSize uint64) *Queue {
	return &Queue{
		content: &ringBuffer{
			buffer: make([]interface{}, initialSize),
			head:   0,
			tail:   0,
			mod:    initialSize,
		},
		len: 0,
	}
}

func (q *Queue) Push(item interface{}) {
	q.lock.Lock()
	c := q.content
	c.tail = ((c.tail + 1) % c.mod)
	if c.tail == c.head {
		var fillFactor uint64 = 10
		//we need to resize

		newLen := c.mod * fillFactor
		newBuff := make([]interface{}, newLen)

		for i := uint64(0); i < c.mod; i++ {
			buffIndex := (c.tail + i) % c.mod
			newBuff[i] = c.buffer[buffIndex]
		}
		//set the new buffer and reset head and tail
		newContent := &ringBuffer{
			buffer: newBuff,
			head:   0,
			tail:   c.mod,
			mod:    c.mod * fillFactor,
		}
		q.content = newContent
	}
	q.len++
	q.content.buffer[q.content.tail] = item
	q.lock.Unlock()
}

func (q *Queue) Length() uint64 {
	res := atomic.LoadUint64(&q.len)
	return res
}

func (q *Queue) Empty() bool {
	return q.Length() == 0
}

//single consumer
func (q *Queue) Pop() (interface{}, bool) {

	if q.Empty() {
		return nil, false
	}
	q.lock.Lock()
	c := q.content
	c.head = ((c.head + 1) % c.mod)
	q.len--
	res := c.buffer[c.head]
	q.lock.Unlock()
	return res, true
}

func (q *Queue) PopMany(count uint64) ([]interface{}, bool) {

	if q.Empty() {
		q.lock.Unlock()
		return nil, false
	}

	q.lock.Lock()
	c := q.content

	if count >= q.len {
		count = q.len
	}

	buffer := make([]interface{}, count)
	for i := uint64(0); i < count; i++ {
		buffer[i] = c.buffer[(c.head+1+i)%c.mod]
	}
	c.head = (c.head + count) % c.mod
	q.len -= count
	q.lock.Unlock()
	return buffer, true
}
