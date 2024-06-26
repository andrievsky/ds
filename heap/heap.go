package heap

import "cmp"

type Heap[T cmp.Ordered] interface {
	Push(value T)
	Peek() T
	Pop() T
	Empty() bool
}
