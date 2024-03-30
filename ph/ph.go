package ph

import "cmp"

type PairingHeap[Elem cmp.Ordered] struct {
	root *PairingTree[Elem]
}

func New[Elem cmp.Ordered]() *PairingHeap[Elem] {
	return &PairingHeap[Elem]{nil}
}

func (heap *PairingHeap[Elem]) Insert(e Elem) {
	heap.root = Meld(heap.root, &PairingTree[Elem]{e, nil})
}

func (heap *PairingHeap[Elem]) Peek() Elem {
	return heap.root.head
}

func (heap *PairingHeap[Elem]) Pop() Elem {
	e := heap.root.head
	heap.root = mergePairs(heap.root.tail)
	return e
}

func (heap *PairingHeap[Elem]) IsEmpty() bool {
	return heap.root == nil
}
