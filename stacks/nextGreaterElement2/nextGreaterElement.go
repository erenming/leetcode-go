package nextGreaterElement2

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

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	stack := New(make([]int, 0))
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	for i := 0; i < 2*n; i++ {
		num := nums[i % n]
		for !stack.empty() && nums[stack.peek() % n] < num {
			res[stack.pop() % n] = num
		}
		if i < n {
			stack.push(i)
		}
	}
	return res
}
