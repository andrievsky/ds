package heap

import "cmp"

type PairingTree[Elem cmp.Ordered] struct {
	head Elem
	tail []*PairingTree[Elem]
}

func Meld[Elem cmp.Ordered](tree1 *PairingTree[Elem], tree2 *PairingTree[Elem]) *PairingTree[Elem] {
	if tree1 == nil {
		return tree2
	}
	if tree2 == nil {
		return tree1
	}
	if tree2.head > tree1.head {
		tree2.tail = append(tree2.tail, tree1)
		return tree2
	}
	tree1.tail = append(tree1.tail, tree2)
	return tree1
}

func IsEmpty[Elem cmp.Ordered](tree *PairingTree[Elem]) bool {
	return tree == nil
}

func mergePairs[Elem cmp.Ordered](trees []*PairingTree[Elem]) *PairingTree[Elem] {
	if len(trees) == 0 {
		return nil
	}
	if len(trees) == 1 {
		return trees[0]
	}
	return Meld(Meld(trees[0], trees[1]), mergePairs(trees[2:]))
}

type PairingHeap[Elem cmp.Ordered] struct {
	root *PairingTree[Elem]
}

func NewMaxHeap[Elem cmp.Ordered]() *PairingHeap[Elem] {
	return &PairingHeap[Elem]{nil}
}

func (heap *PairingHeap[Elem]) Push(e Elem) {
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
