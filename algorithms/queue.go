package algorithms

type QueueNode[T any] struct {
	Val  T
	Next *QueueNode[T]
}

type Queue[T any] struct {
	Size int
	Head *QueueNode[T]
	Tail *QueueNode[T]
}

func (q *Queue[T]) Enqueue(val T) {
	q.Size++
	node := &QueueNode[T]{Val: val}
	if q.Head == nil {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		q.Tail = node
	}
}

func (q *Queue[T]) Dequeue() T {
	if q.Head == nil {
		var zero T
		return zero
	}

	q.Size--
	head := q.Head
	q.Head = head.Next

	head.Next = nil // free

	return head.Val
}

func (q *Queue[T]) Peek() *T {
	if q.Head == nil {
		return nil
	}
	return &q.Head.Val
}
