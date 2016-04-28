package trees

import (
//"fmt"
)

type BinaryHeap struct {
	tree []int
}

func NewBinaryHeap() *BinaryHeap {
	var heap = new(BinaryHeap)
	heap.tree = make([]int, 0)
	return heap
}

func (this *BinaryHeap) Add(node int) {
	var i = len(this.tree)
	this.tree = append(this.tree, node)
	var j = parent(i)
	for j != -1 && this.tree[i] < this.tree[j] {
		this.tree[i], this.tree[j] = this.tree[j], this.tree[i]
		i = j
		j = parent(i)
	}
}

func (this *BinaryHeap) Next() bool {
	return len(this.tree) > 0
}

func (this *BinaryHeap) Extract() int {
	var n = len(this.tree)
	if n == 0 {
		return -1
	}
	var next = this.tree[0]
	this.tree = append(this.tree[:0], this.tree[1:]...)
	this.heapify(n - 1)
	return next
}

func (this *BinaryHeap) heapify(i int) {
	var n = len(this.tree)
	if n == 0 {
		return
	}
	var m = i
	var p = parent(i - 1)
	if p < 0 {
		return
	}
	var l = left(p)
	if l <= n && this.tree[i-1] < this.tree[l] {
		m = l
	}
	var r = right(p)
	if r <= n && this.tree[i-1] < this.tree[r] {
		m = r
	}
	if m != i {
		this.tree[i-1], this.tree[m] = this.tree[m], this.tree[i-1]
		this.heapify(m)
	}
}

func (this *BinaryHeap) ToArray() []int {
	return this.tree
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}

func parent(i int) int {
	if i == 0 {
		return -1
	}
	return (i - 1) / 2
}
