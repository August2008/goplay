package trees

import (
//"fmt"
)

type BHeap struct {
	tree []int
}

func NewBHeap() *BHeap {
	var heap = new(BHeap)
	heap.tree = make([]int, 0)
	return heap
}

func (this *BHeap) Add(node int) {
	var i = len(this.tree)
	this.tree = append(this.tree, node)
	var j = parent(i + 1)
	for j != -1 && this.tree[i] < this.tree[j] {
		this.tree[i], this.tree[j] = this.tree[j], this.tree[i]
		i = j
		j = parent(i)
	}
}

func (this *BHeap) Next() bool {
	return len(this.tree) > 0
}

func (this *BHeap) Extract() int {
	var n = len(this.tree)
	if n == 0 {
		return -1
	}
	var next = this.tree[0]
	this.tree = append(this.tree[:0], this.tree[1:]...)
	this.heapify(n - 1)
	return next
}

func (this *BHeap) heapify(i int) {
	var n = len(this.tree)
	if n == 0 {
		return
	}
	var m, j = i - 1, i - 1
	var p = parent(j)
	if p < 0 {
		return
	}
	var l = left(p)
	if l < n && this.tree[j] < this.tree[l] {
		m = l
	}
	var r = right(p)
	if r < n && this.tree[j] < this.tree[r] {
		m = r
	}
	if m != j {
		this.tree[j], this.tree[m] = this.tree[m], this.tree[j]
		this.heapify(m)
	}
}

func (this *BHeap) ToArray() []int {
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
