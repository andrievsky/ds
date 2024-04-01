package binary

type Node struct {
	key      string
	priority int
	index    int
}

type PriorityQueue struct {
	nodes map[string]*Node
	queue []*Node
}

func New() *PriorityQueue {
	return &PriorityQueue{make(map[string]*Node), make([]*Node, 0)}
}

func (t *PriorityQueue) Push(key string, prioruty int) {
	node := &Node{key: key, priority: prioruty, index: len(t.queue)}
	t.nodes[key] = node
	t.queue = append(t.queue, node)
	up(t.queue, len(t.queue)-1)
}

func (t *PriorityQueue) Peek() string {
	return t.queue[0].key
}

func (t *PriorityQueue) Pop() string {
	node := t.queue[0]
	delete(t.nodes, node.key)
	last := len(t.queue) - 1
	t.queue[0] = t.queue[last]
	t.queue[0].index = 0
	t.queue = t.queue[:last]
	down(t.queue, 0)
	return node.key
}

func (t *PriorityQueue) Remove(key string) string {
	node := t.nodes[key]
	last := len(t.queue) - 1
	i := node.index
	if last != node.index {
		t.queue[i] = t.queue[last]
		t.queue[i].index = 0
		t.queue = t.queue[:last]
		if !down(t.queue, i) {
			up(t.queue, i)
		}
	}
	return node.key
}

func (t *PriorityQueue) Update(key string, priority int) {
	node := t.nodes[key]
	node.priority = priority
	if !down(t.queue, node.index) {
		up(t.queue, node.index)
	}
}

func (t *PriorityQueue) Empty() bool {
	return len(t.queue) == 0
}

func up(queue []*Node, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || queue[j].priority >= queue[i].priority {
			break
		}
		queue[i], queue[j] = queue[j], queue[i]
		queue[i].index = i
		queue[j].index = j
		j = i
	}
}

func down(queue []*Node, i0 int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= len(queue) || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < len(queue) && queue[j2].priority < queue[j1].priority {
			j = j2 // = 2*i + 2  // right child
		}
		if queue[j].priority >= queue[i].priority {
			break
		}
		queue[i], queue[j] = queue[j], queue[i]
		queue[i].index = i
		queue[j].index = j
		i = j
	}
	return i > i0
}
