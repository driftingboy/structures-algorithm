package queue

// 避免数据数据移动, 会牺牲一个空间
type RingQueue struct {
	head int
	tail int

	size int
	data []interface{}
}

func NewRingQueue(size int) *RingQueue {
	return &RingQueue{
		size: size,
		data: make([]interface{}, 0, size),
	}
}

func (q *RingQueue) Push(v interface{}) bool {
	if (q.tail+1)%q.size == q.head {
		return false
	}
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.size
	return true
}

func (q *RingQueue) Pop() (v interface{}, ok bool) {
	if v, ok = q.Head(); ok {
		q.head = (q.head + 1) % q.size
		return
	}

	return nil, false
}

func (q *RingQueue) Head() (v interface{}, ok bool) {
	if q.head == q.tail {
		return nil, false
	}
	return q.data[q.head], true
}
