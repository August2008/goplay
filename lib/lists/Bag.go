package lists

type Bag struct {
	size  int
	first *node
}

type BagIterator struct {
	bag     *Bag
	last    *node
	Current interface{}
}

func NewBag() *Bag {
	return new(Bag)
}

func (b *Bag) Add(item interface{}) {
	var first = b.first
	b.first = new(node)
	b.first.item = item
	b.first.next = first
	b.size++
}

func (b *Bag) Interator() *BagIterator {
	return &BagIterator{bag: b, last: b.first}
}

func (it *BagIterator) Next() bool {
	if it.last != nil {
		it.Current = it.last.item
		it.last = it.last.next
		return true
	}
	it.Current = nil
	return false
}
