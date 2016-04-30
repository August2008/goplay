package trees

//https://en.wikipedia.org/wiki/Binary_heap

import (
//"fmt"
)

type Interface interface {
	Less(i, j int) bool
}

type BHeap struct {
	tree []Interface
}

func NewBHeap(n int) *BHeap {
	var h = new(BHeap)
	h.tree = make([]Interface, 0, n)
	return h
}

func (h *BHeap) Push(node Interface) {
	var i = len(h.tree)
	h.tree = append(h.tree, node) //always add at the end
	var j = parent(i)
	for j != -1 && node.Less(i, j) { //maintain transitive relation
		//whenever A < B and B < C, then also A < C
		h.swap(i, j) //replace parent with it
		i = j
		j = parent(i) //keep bubbling it up until it's less than parent
	}
}

func (h *BHeap) Pop() interface{} {
	var n = len(h.tree)
	if n == 0 {
		return -1
	}
	var next = h.next() //get next and remove it + move last to first
	h.heapify(0)
	return next
}

func (h *BHeap) Empty() bool {
	return len(h.tree) == 0
}

func (h *BHeap) heapify(i int) {
	var n = len(h.tree) - 1
	if n == 0 {
		return //just one left
	}
	var j = i
	var l = left(i)
	if l <= n && h.tree[l].Less(l, i) {
		j = l
	}
	var r = right(i)
	if r <= n && h.tree[r].Less(r, j) {
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

func (h *BHeap) next() interface{} {
	var next = h.tree[0]                       //first element is always min/max
	h.tree = append(h.tree[:0], h.tree[1:]...) //remove it from the tree
	var n = len(h.tree) - 1
	if n > 2 {
		h.tree = append(h.tree[n:], h.tree[:n]...) //move last to first
	}
	return next
}

func (h *BHeap) ToArray() []Interface {
	return h.tree
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
