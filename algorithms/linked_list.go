package algorithms

type SingleLinkedList struct {
	Val  int
	Next *SingleLinkedList
}

func (sll *SingleLinkedList) Append(val int) {
	for n := sll; ; n = n.Next {
		if n.Next == nil {
			n.Next = &SingleLinkedList{Val: val}
			break
		}
	}
}

type DoubleLinkedList struct {
	Prev *DoubleLinkedList
	Val  int
	Next *DoubleLinkedList
}

func (dll *DoubleLinkedList) Append(val int) {
	for n := dll; ; n = n.Next {
		if n.Next == nil {
			n.Next = &DoubleLinkedList{Prev: n, Val: val}
			break
		}
	}
}

type CircularLinkedList struct {
	Val  int
	Next *CircularLinkedList
}

func (cll *CircularLinkedList) Append(val int) {
	for n := cll; ; n = n.Next {
		if n.Next == nil {
			n.Next = &CircularLinkedList{Val: val}
			break
		}
	}
}
