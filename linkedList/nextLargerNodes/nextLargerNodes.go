package nextLargerNodes

type ListNode struct {
	Val  int
	Next *ListNode
}

//func nextLargerNodes(head *ListNode) []int {
//	if head == nil {
//		return nil
//	}
//	res := make([]int, 0)
//	pre, next := head, head
//	for pre != nil {
//		for next != nil {
//			if next.Val > pre.Val {
//				res = append(res, next.Val)
//				break
//			}
//			next = next.Next
//		}
//		if next == nil {
//			res = append(res, 0)
//		}
//		pre = pre.Next
//		next = pre
//	}
//	return res
//}

// stack version
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

func nextLargerNodes(head *ListNode) []int {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	res := make([]int, len(nums))
	// top to down, increasing stack
	stack := &intStack{st: make([]int, 0, len(nums))}
	for i := len(nums)-1; i >= 0; i-- {
		for !stack.empty() && stack.peek() <= nums[i] {
			stack.pop()
		}
		if stack.empty() {
			res[i] = 0
		} else {
			res[i] = stack.peek()
		}
		stack.push(nums[i])
	}
	return res
}