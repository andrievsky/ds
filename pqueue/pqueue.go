package pqueue

import "cmp"

type PriorityQueue[Key comparable, Priority cmp.Ordered] interface {
	Push(key Key, priority Priority)
	Peek() Key
	Pop() Key
	Update(key Key, priority Priority)
	IsEmpty() bool
}
