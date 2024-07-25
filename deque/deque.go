package deque

type Deque[T comparable] struct {
	head *node[T]
	tail *node[T]
}

type node[T comparable] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

func New[T comparable]() *Deque[T] {
	return &Deque[T]{}
}

func (t *Deque[T]) PushFront(value T) {
	node := &node[T]{value, t.head, nil}
	if t.head == nil {
		t.head = node
		t.tail = node
		return
	}
	node.next = t.head
	t.head.prev = node
	t.head = node
}

func (t *Deque[T]) PushBack(value T) {
	node := &node[T]{value, nil, t.tail}
	if t.tail == nil {
		t.head = node
		t.tail = node
		return
	}
	node.prev = t.tail
	t.tail.next = node
	t.tail = node
}

func (t *Deque[T]) PopFront() T {
	value := t.head.value
	t.head = t.head.next
	if t.head == nil {
		t.tail = nil
	} else {
		t.head.prev = nil
	}
	return value
}

func (t *Deque[T]) PopBack() T {
	value := t.tail.value
	t.tail = t.tail.prev
	if t.tail == nil {
		t.head = nil
	} else {
		t.tail.next = nil
	}
	return value
}

func (t *Deque[T]) Empty() bool {
	return t.head == nil
}
