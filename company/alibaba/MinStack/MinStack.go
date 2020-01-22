package MinStack

type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 || x <= this.GetMin() {
		this.min = append(this.min, x)
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	if this.Top() == this.GetMin() {
		this.min = append([]int{}, this.min[:len(this.min)-1]...)
	}
	this.stack = append([]int{}, this.stack[:len(this.stack)-1]...)
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
