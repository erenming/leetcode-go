package isValid

type runeStack struct {
	st []rune
}

func New() *runeStack {
	return &runeStack{st: make([]rune, 0)}
}

func (s *runeStack) Empty() bool {
	return len(s.st) == 0
}

func (s *runeStack) Push(n rune) {
	s.st = append(s.st, n)
}

func (s *runeStack) Pop() rune {
	if s.Empty() {
		panic("stack is Empty")
	}
	i := s.st[len(s.st)-1]
	tmp := s.st[:len(s.st)-1]
	s.st = make([]rune, len(s.st)-1)
	copy(s.st, tmp)
	return i
}

func (s *runeStack) Peek() rune {
	if s.Empty() {
		panic("stack is Empty")
	}
	return s.st[len(s.st)-1]
}

func isValid(s string) bool {
	stack := New()
	source := []rune(s)
	n := len(source)
	for i := 0; i < n; i++ {
		c := source[i]
		if stack.Empty() {
			stack.Push(c)
			continue
		}
		top := stack.Peek()
		if (top == '(' && c == ')') || (top == '[' && c == ']') || (top == '{' && c == '}') {
			stack.Pop()
		} else {
			stack.Push(c)
		}
	}
	return stack.Empty()
}
