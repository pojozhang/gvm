package runtime

type Stack struct {
	maxSize uint
	size    uint
	head    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (s *Stack) push(frame *Frame) {
	if s.size >= s.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if s.head != nil {
		s.head.next = frame
	}
	s.head = frame
	s.size++
}

func (s *Stack) pop() *Frame {
	peak := s.head
	s.head = s.head.next
	s.size--
	return peak
}

func (s *Stack) top() *Frame {
	return s.head
}
