package stacks

type intStack struct {
	st []int
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

type MinStack struct {
	mins intStack
	nums intStack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.nums.push(x)
	if this.mins.empty() || x <= this.mins.peek() {
		this.mins.push(x)
	}
}

func (this *MinStack) Pop() {
	x := this.nums.pop()
	if x == this.mins.peek() {
		this.mins.pop()
	}
}

func (this *MinStack) Top() int {
	return this.nums.peek()
}

func (this *MinStack) GetMin() int {
	return this.mins.peek()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
