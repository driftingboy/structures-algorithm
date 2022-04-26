package queue

// 有界队列
type ArrayQueue struct {
	head int
	tail int

	size int
	data []interface{}
}

func NewArrayQueue(size int) *ArrayQueue {
	return &ArrayQueue{
		size: size,
		data: make([]interface{}, 0, size),
	}
}

func (q *ArrayQueue) Push(v interface{}) bool {
	if q.tail == q.size {
		if q.head == 0 {
			return false
		}
		// 空间不足且有头部空闲空间，再统一移动
		for i, v := range q.data[q.head:] {
			q.data[i] = v
			q.tail -= q.head
		}
	}
	q.data[q.tail] = v
	q.tail++
	return true
}

func (q *ArrayQueue) Pop() (v interface{}, ok bool) {
	if v, ok = q.Head(); ok {
		q.head++
		return
	}

	return nil, false
}

func (q *ArrayQueue) Head() (v interface{}, ok bool) {
	if q.head == q.tail {
		return nil, false
	}
	return q.data[q.head], true
}
