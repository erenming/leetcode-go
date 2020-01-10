package mergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{0, nil}
	root := r
	for l1 != nil || l2 != nil {
		if l1 == nil {
			r.Next = l2
			l2 = nil
		} else if l2 == nil {
			r.Next = l1
			l1 = nil
		} else {
			if l1.Val > l2.Val {
				r.Next = l2
				l2 = l2.Next
				r = r.Next
			} else {
				r.Next = l1
				l1 = l1.Next
				r = r.Next
			}
		}
	}
	return root.Next
}
