package gods

// Queue is a FIFO (first in first out) data structure implementation
// It is based on a deque container and focuses its API on core
// functionalities: Enqueue, Dequeue, Head, Size, Empty. Every operations time complexity
// is O(1).
type Queue struct {
	*Deque
}

func NewQueue() *Queue {
	return &Queue{NewDeque()}
}

// Enqueue adds an item at the back of the queue
func (q *Queue) Enqueue(item interface{}) {
	q.Prepend(item)
}

// Dequeue removes and returns the front queue item
func (q *Queue) Dequeue() interface{} {
	return q.Pop()
}

// Head returns the front queue item
func (q *Queue) Head() interface{} {
	return q.Last()
}
