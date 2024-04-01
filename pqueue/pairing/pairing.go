package pairing

import (
	"cmp"
	"fmt"
	"strings"
)

type LinkedList[T comparable] struct {
	value T
	next  *LinkedList[T]
}

type Node[Key comparable, Priority cmp.Ordered] struct {
	parent   *Node[Key, Priority]
	key      Key
	priority Priority
	nodes    *LinkedList[*Node[Key, Priority]]
}

type PriorityQueue[Key comparable, Priority cmp.Ordered] struct {
	root  *Node[Key, Priority]
	nodes map[Key]*Node[Key, Priority]
}

func New[Key comparable, Priority cmp.Ordered]() *PriorityQueue[Key, Priority] {
	return &PriorityQueue[Key, Priority]{nil, make(map[Key]*Node[Key, Priority])}
}

func (t *PriorityQueue[Key, Priority]) Push(key Key, priority Priority) {
	node := &Node[Key, Priority]{nil, key, priority, nil}
	t.root = meld(t.root, node)
	t.nodes[key] = node
}

func (t *PriorityQueue[Key, Priority]) Peek() Key {
	return t.root.key
}

func (t *PriorityQueue[Key, Priority]) Pop() Key {
	key := t.root.key
	delete(t.nodes, key)
	t.root = mergePairs(t.root.nodes)
	if t.root != nil {
		t.root.parent = nil
	}
	return key
}

func (t *PriorityQueue[Key, Priority]) Update(key Key, priority Priority) {
	node := t.nodes[key]
	if priority == node.priority {
		return
	}
	parent := node.parent
	if priority > node.priority {
		node.priority = priority
		if parent != nil && priority > parent.priority {
			parent.nodes = remove(parent.nodes, node)
			node.parent = nil
			if t.root != node {
				t.root = meld(t.root, node)
			}
		}
		return
	}

	node.priority = priority
	if node.nodes == nil {
		return
	}
	childNode := mergePairs(node.nodes)
	if parent != nil {
		replace(parent.nodes, node, childNode)
		childNode.parent = parent
	} else {
		t.root = childNode
		childNode.parent = nil
	}
	node.parent = nil
	node.nodes = nil

	t.root = meld(t.root, node)
}

func (t *PriorityQueue[Key, Priority]) Empty() bool {
	return t.root == nil
}

func (t *PriorityQueue[Key, Priority]) Pretty() string {
	sb := strings.Builder{}
	var pretty func(node *Node[Key, Priority], depth int)
	pretty = func(node *Node[Key, Priority], depth int) {
		if node == nil {
			return
		}
		sb.WriteString(strings.Repeat("  ", depth))
		sb.WriteString(fmt.Sprintf("%v(%v, %v)\n", node.key, node.priority, node.parent))
		for child := node.nodes; child != nil; child = child.next {
			pretty(child.value, depth+1)
		}
	}
	pretty(t.root, 0)
	return sb.String()
}

func meld[Key comparable, Priority cmp.Ordered](node1 *Node[Key, Priority], node2 *Node[Key, Priority]) *Node[Key, Priority] {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	if node2.priority > node1.priority {
		node1.parent = node2
		node2.nodes = add(node2.nodes, node1)
		return node2
	}
	node2.parent = node1
	node1.nodes = add(node1.nodes, node2)
	return node1
}

func mergePairs[Key comparable, Priority cmp.Ordered](nodes *LinkedList[*Node[Key, Priority]]) *Node[Key, Priority] {
	if nodes == nil {
		return nil
	}
	if nodes.next == nil {
		return nodes.value
	}
	return meld(meld(nodes.value, nodes.next.value), mergePairs(nodes.next.next))
}

func add[T comparable](list *LinkedList[T], value T) *LinkedList[T] {
	return &LinkedList[T]{value, list}
}

func remove[T comparable](list *LinkedList[T], value T) *LinkedList[T] {
	if list == nil {
		return nil
	}
	if list.value == value {
		return list.next
	}
	parent := list
	for parent.next != nil {
		if parent.next.value == value {
			parent.next = parent.next.next
			return list
		}
		parent = parent.next
	}
	return list
}

func replace[T comparable](list *LinkedList[T], value T, newValue T) {
	if list == nil {
		return
	}
	for node := list; node != nil; node = node.next {
		if node.value == value {
			node.value = newValue
		}
	}
}
