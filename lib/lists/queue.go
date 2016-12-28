package lists

type Queue struct {
	size  int
	last  *node
	first *node
}

func NewQueue() *Queue {
	return new(Queue)
}

func (q *Queue) Enqueue(item interface{}) {
	var last = q.last
	q.last = new(node)
	q.last.item = item
	if q.size == 0 {
		q.first = q.last
	} else {
		last.next = q.last
	}
	q.size++
}

func (q *Queue) Dequeue() interface{} {
	var item interface{} = nil
	if q.size > 0 {
		item = q.first.item
		q.first = q.first.next
		q.size--
	} else {
		q.last = nil
		q.first = nil
	}
	return item
}

func (q *Queue) Peek() interface{} {
	if q.size > 0 {
		return q.first.item
	}
	return nil
}

func (q *Queue) Size() int {
	return q.size
}
