package stacks

import "testing"

func TestIsValid(t *testing.T) {
	isValid("()[]{}")
}

func TestMinStack_GetMin(t *testing.T) {
	obj := Constructor()
	obj.Push(0)
	obj.Push(1)
	obj.Push(0)
	obj.GetMin()
	obj.Pop()
	obj.GetMin()
}
