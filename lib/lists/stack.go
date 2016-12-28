package lists

type Stack struct {
	size  int
	first *node
}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Push(item interface{}) {
	var prev = s.first
	s.first = new(node)
	s.first.item = item
	s.first.next = prev
	s.size++
}

func (s *Stack) Pop() interface{} {
	var item interface{} = nil
	if s.size > 0 {
		item = s.first.item
		s.first = s.first.next
		s.size--
	} else {
		s.first = nil
	}
	return item
}

func (s *Stack) Peek() interface{} {
	var item interface{} = nil
	if s.size > 0 {
		item = s.first.item
	}
	return item
}

func (s *Stack) Size() int {
	return s.size
}
