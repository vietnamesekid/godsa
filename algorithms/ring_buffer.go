package algorithms

import "errors"

type RingBuffer[T any] struct {
	Buffer []T
	Head   int
	Tail   int
	Size   int
	Cap    int
}

func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	return &RingBuffer[T]{
		Buffer: make([]T, capacity),
		Cap:    capacity,
	}
}

// IsFull checks if the ring buffer is full.
func (rb *RingBuffer[T]) IsFull() bool {
	return rb.Size == rb.Cap
}

// IsEmpty checks if the ring buffer is empty.
func (rb *RingBuffer[T]) IsEmpty() bool {
	return rb.Size == 0
}

// Write adds a new element to the ring buffer.
func (rb *RingBuffer[T]) Write(data T) error {
	if rb.IsFull() {
		return errors.New("buffer is full")
	}

	rb.Buffer[rb.Head] = data
	rb.Head = (rb.Head + 1) % rb.Cap
	rb.Size++

	return nil
}

// Read retrieves and removes the oldest element from the ring buffer.
func (rb *RingBuffer[T]) Read() (T, error) {
	var zero T
	if rb.IsEmpty() {
		return zero, errors.New("buffer is empty")
	}

	data := rb.Buffer[rb.Tail]
	rb.Tail = (rb.Tail + 1) % rb.Cap
	rb.Size--

	return data, nil
}

// Peek retrieves the oldest element from the ring buffer without removing it.
func (rb *RingBuffer[T]) Peek() (T, error) {
	var zero T
	if rb.IsEmpty() {
		return zero, errors.New("buffer is empty")
	}

	return rb.Buffer[rb.Tail], nil
}
