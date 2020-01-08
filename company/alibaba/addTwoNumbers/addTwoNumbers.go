package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	next := &ListNode{0, nil}
	add := 0
	root := next
	for l1 != nil || l2 != nil {
		if l1 == nil {
			next.Next = &ListNode{(l2.Val + add) % 10, nil}
			add = (l2.Val + add) / 10
			l2 = l2.Next
		} else if l2 == nil {
			next.Next = &ListNode{(l1.Val + add) % 10, nil}
			add = (l1.Val + add) / 10
			l1 = l1.Next
		} else {
			tmp := l1.Val + l2.Val + add
			next.Next = &ListNode{tmp % 10, nil}
			add = tmp / 10
			l1 = l1.Next
			l2 = l2.Next
		}
		next = next.Next
	}
	if add == 1 {
		next.Next = &ListNode{1, nil}
	}
	return root.Next
}
