package dailyTemperatures

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

func dailyTemperatures(T []int) []int {
	stack := New(make([]int, 0))
	res := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		if stack.empty() {
			stack.push(i)
		} else if T[i] <= T[stack.peek()] {
			stack.push(i)
		} else {
			for !stack.empty() && (T[i] > T[stack.peek()]) {
				cur := stack.pop()
				res[cur] = i - cur
			}
			stack.push(i)
		}
	}
	return res
}
