package pkg

type Stack struct {
	buf []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Iter() []int {
	return s.buf
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Len() int {
	return len(s.buf)
}

func (s *Stack) Push(i int) {
	s.buf = append(s.buf, i)
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.buf[s.Len()-1]
}

func (s *Stack) Pop() (i int) {
	i = s.Peek()
	s.buf = s.buf[0 : s.Len()-1]
	return
}
