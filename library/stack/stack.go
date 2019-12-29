package stack

type intStack struct {
	st []int
}

func New(data []int) *intStack {
	return &intStack{st: data}
}

func (s *intStack) empty() bool {
	return len(s.st) == 0
}

func (s *intStack) push(n int) {
	s.st = append(s.st, n)
}

func (s *intStack) pop() int {
	if s.empty() {
		panic("stack is empty")
	}
	i := s.st[len(s.st)-1]
	s.st = s.st[:len(s.st)-1]
	return i
}

func (s *intStack) peek() int {
	if s.empty() {
		panic("stack is empty")
	}
	return s.st[len(s.st)-1]
}
