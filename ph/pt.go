package ph

import "cmp"

type PairingTree[Elem cmp.Ordered] struct {
	head Elem
	tail []*PairingTree[Elem]
}

func Insert[Elem cmp.Ordered](heap *PairingTree[Elem], e Elem) *PairingTree[Elem] {
	return Meld(heap, &PairingTree[Elem]{e, nil})
}

func Peek[Elem cmp.Ordered](tree *PairingTree[Elem]) Elem {
	return tree.head
}

func Pop[Elem cmp.Ordered](tree1 *PairingTree[Elem]) (Elem, *PairingTree[Elem]) {
	return tree1.head, mergePairs(tree1.tail)
}

func Meld[Elem cmp.Ordered](tree1 *PairingTree[Elem], tree2 *PairingTree[Elem]) *PairingTree[Elem] {
	if tree1 == nil {
		return &PairingTree[Elem]{tree2.head, tree2.tail}
	}
	if tree2 == nil {
		return &PairingTree[Elem]{tree1.head, tree1.tail}
	}
	if tree2.head > tree1.head {
		return &PairingTree[Elem]{tree2.head, append(tree2.tail, tree1)}
	}
	return &PairingTree[Elem]{tree1.head, append(tree1.tail, tree2)}
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
