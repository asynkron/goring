package goring

import "sync"

type ringBuffer struct {
	buffer []interface{}
	head   int
	tail   int
	mod    int
}

type Queue struct {
	content *ringBuffer
	len     int
	lock    sync.RWMutex
}

func New(initialSize int) *Queue {
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
		fillFactor := 50
		//we need to resize

		newLen := c.mod * fillFactor
		newBuff := make([]interface{}, newLen)

		for i := 0; i < c.mod; i++ {
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

func (q *Queue) Length() int {
	q.lock.RLock()
	len := q.len
	q.lock.RUnlock()
	return len
}

func (q *Queue) Empty() bool {
	return q.Length() == 0
}

func (q *Queue) Pop() (interface{}, bool) {
	q.lock.Lock()

	if q.len == 0 {

		q.lock.Unlock()
		return nil, false
	}
	c := q.content
	c.head = ((c.head + 1) % c.mod)
	q.len--
	q.lock.Unlock()
	return c.buffer[c.head], true
}

func (q *Queue) PopMany(count int) ([]interface{}, bool) {
	q.lock.Lock()
	if q.len == 0 {
		q.lock.Unlock()
		return nil, false
	}
	c := q.content

	if count >= q.len {
		count = q.len
	}

	buffer := make([]interface{}, count)
	for i := 0; i < count; i++ {
		buffer[i] = c.buffer[(c.head+1+i)%c.mod]
	}
	c.head = (c.head + count) % c.mod
	q.len -= count
	q.lock.Unlock()
	return buffer, true
}
