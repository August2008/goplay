package trees

//https://en.wikipedia.org/wiki/Binary_tree
/*
Binary heap with n elements has (log n) height thus add/remove are O(log n)
*/

import "fmt"

type node struct {
	value    interface{}
	priority float64
}

type BHeap struct {
	set  map[interface{}]*node //used in update to find the node in O(1)
	tree []*node
	less func(i, j int) bool
}

func NewMinHeap() *BHeap {
	var h = new(BHeap)
	h.init(func(i, j int) bool { //less func to decide the priority
		return h.tree[i].priority < h.tree[j].priority
	})
	return h
}

func NewMaxHeap() *BHeap {
	var h = new(BHeap)
	h.init(func(i, j int) bool {
		return h.tree[i].priority > h.tree[j].priority
	})
	return h
}

func (h *BHeap) init(less func(i, j int) bool) {
	h.set = make(map[interface{}]*node)
	h.tree = make([]*node, 0)
	h.less = less
}

func (h *BHeap) Push(obj interface{}, priority float64) {
	if _, ok := h.set[obj]; ok {
		return
	}
	var (
		i    = h.Len()
		node = &node{value: obj, priority: priority}
	)
	h.tree = append(h.tree, node)
	for j := parent(i); j != -1 && h.less(i, j); { //maintain transitive relation
		h.swap(i, j) //swap with parent
		i = j
		j = parent(i) //keep swapping until it's less than parent
	}
	h.set[obj] = node
}

func (h *BHeap) Peek() interface{} {
	return h.first(false).value
}

func (h *BHeap) Pop() interface{} {
	if n := h.Len(); n == 0 {
		return nil
	}
	var next = h.next() //get next and remove it + move last to first
	delete(h.set, next.value)
	h.heapify(0)
	return next.value
}

func (h *BHeap) Update(obj interface{}, priority float64) {
	if node, ok := h.set[obj]; ok {
		node.priority = priority
	}
	h.heapify(0)
}

func (h *BHeap) Len() int {
	return len(h.tree)
}

func (h *BHeap) heapify(i int) {
	var n = h.Len() - 1
	if n == 0 {
		return
	}
	var j = i
	if l := left(i); l <= n && h.less(l, i) {
		j = l
	}
	if r := right(i); r <= n && h.less(r, j) {
		j = r
	}
	if j != i {
		h.swap(i, j)
		h.heapify(j)
	}
}

func (h *BHeap) swap(i, j int) {
	h.tree[i], h.tree[j] = h.tree[j], h.tree[i]
}

func (h *BHeap) next() node {
	var next = h.first(true) //first element is always next
	if n := h.Len() - 1; n > 2 {
		h.tree = append(h.tree[n:], h.tree[:n]...) //move last to first
	}
	return next
}

func (h *BHeap) first(remove bool) node {
	var next = h.tree[0]
	if remove {
		h.tree = append(h.tree[:0], h.tree[1:]...) //remove it from the tree
	}
	return *next
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func parent(i int) int {
	if i == 0 {
		return -1
	}
	return (i - 1) / 2
}

func (h *BHeap) Debug() {
	for _, v := range h.tree {
		fmt.Printf("%v, %v\n", v.value, v.priority)
	}
}
