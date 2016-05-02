package trees

//https://en.wikipedia.org/wiki/Binary_tree

type Node struct {
	Value interface{}
	Ordinal int
}

type BHeap struct {
	tree []Node
	less func (i, j int) bool
}

func NewMinHeap(n int) *BHeap {
	var h = new(BHeap)
	h.tree = make([]Node, 0, n)
	h.less = func (i, j int) bool {
		return h.tree[i].Ordinal < h.tree[j].Ordinal
	}
	return h
}

func NewMaxHeap(n int) *BHeap {
	var h = new(BHeap)
	h.tree = make([]Node, 0, n)
	h.less = func (i, j int) bool {
		return h.tree[i].Ordinal > h.tree[j].Ordinal
	}	
	return h
}

func (h *BHeap) Push(node *Node) {
	var i = h.Len()
	h.tree = append(h.tree, *node) 
	for j := parent(i); j != -1 && h.less(i, j); { //maintain transitive relation
		h.swap(i, j) //swap with parent
		i = j
		j = parent(i) //keep swapping until it's less than parent
	}
}

func (h *BHeap) Pop() *Node {
	if n := h.Len(); n == 0 {
		return nil 
	}
	var next = h.next() //get next and remove it + move last to first
	h.heapify(0)
	return &next 
}

func (h *BHeap) Len() int {
	return len(h.tree)
}

func (h *BHeap) heapify(i int) {
	var n = h.Len() - 1
	if  n == 0 {
		return 
	}
	var j = i
	if  l := left(i); l <= n && h.less(l, i) {
		j = l
	}	
	if  r := right(i); r <= n && h.less(r, j) {
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

func (h *BHeap) next() Node {
	var next = h.first() //first element is always next
	if n := h.Len() - 1; n > 2 {
		h.tree = append(h.tree[n:], h.tree[:n]...) //move last to first
	}
	return next
}

func (h *BHeap) first() Node {
	var next = h.tree[0]                       
	h.tree = append(h.tree[:0], h.tree[1:]...) //remove it from the tree
	return next;
}

func (h *BHeap) ToArray() []Node {
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
