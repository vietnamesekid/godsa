package algorithms

type StackNode[T any] struct {
	Val  T
	Prev *StackNode[T]
}

type Stack[T any] struct {
	Size int
	Head *StackNode[T]
}

func (s *Stack[T]) Push(val T) {
	s.Size++
	node := &StackNode[T]{Val: val}

	if s.Head == nil {
		s.Head = node
	} else {
		node.Prev = s.Head
		s.Head = node
	}
}

func (s *Stack[T]) Pop() T {
	if s.Head == nil {
		var zero T
		return zero
	}

	s.Size--
	val := s.Head.Val
	s.Head = s.Head.Prev
	return val
}

func (s *Stack[T]) Peek() *T {
	if s.Head == nil {
		return nil
	}
	return &s.Head.Val
}
