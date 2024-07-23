package binary

type Node[K comparable, V any] struct {
	key      K
	priority V
	index    int
}

type PriorityQueue[K comparable, V any] struct {
	nodes map[K]*Node[K, V]
	queue []*Node[K, V]
	less  func(a, b V) bool
}

func New[K comparable, V any](less func(a, b V) bool) *PriorityQueue[K, V] {
	return &PriorityQueue[K, V]{make(map[K]*Node[K, V]), make([]*Node[K, V], 0), less}
}

func (t *PriorityQueue[K, V]) Push(key K, priority V) {
	node := &Node[K, V]{key: key, priority: priority, index: len(t.queue)}
	t.nodes[key] = node
	t.queue = append(t.queue, node)
	t.up(len(t.queue) - 1)
}

func (t *PriorityQueue[K, V]) Peek() (K, V) {
	return t.queue[0].key, t.queue[0].priority
}

func (t *PriorityQueue[K, V]) Pop() (K, V) {
	node := t.queue[0]
	delete(t.nodes, node.key)
	last := len(t.queue) - 1
	t.queue[0] = t.queue[last]
	t.queue[0].index = 0
	t.queue = t.queue[:last]
	t.down(0)
	return node.key, node.priority
}

func (t *PriorityQueue[K, V]) Delete(key K) {
	node := t.nodes[key]
	last := len(t.queue) - 1
	i := node.index
	if last != node.index {
		t.queue[i] = t.queue[last]
		t.queue[i].index = 0
		t.queue = t.queue[:last]
		if !t.down(i) {
			t.up(i)
		}
	}
}

func (t *PriorityQueue[K, V]) Update(key K, priority V) {
	node := t.nodes[key]
	node.priority = priority
	if !t.down(node.index) {
		t.up(node.index)
	}
}

func (t *PriorityQueue[K, V]) Empty() bool {
	return len(t.queue) == 0
}

func (t *PriorityQueue[K, V]) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !t.less(t.queue[j].priority, t.queue[i].priority) {
			break
		}
		t.queue[i], t.queue[j] = t.queue[j], t.queue[i]
		t.queue[i].index = i
		t.queue[j].index = j
		j = i
	}
}

func (t *PriorityQueue[K, V]) down(i0 int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= len(t.queue) || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < len(t.queue) && t.less(t.queue[j2].priority, t.queue[j1].priority) {
			j = j2
		}
		if !t.less(t.queue[j].priority, t.queue[i].priority) {
			break
		}
		t.queue[i], t.queue[j] = t.queue[j], t.queue[i]
		t.queue[i].index = i
		t.queue[j].index = j
		i = j
	}
	return i > i0
}
