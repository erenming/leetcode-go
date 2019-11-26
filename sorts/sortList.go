package sorts

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := &ListNode{}
	pre.Next = head
	slow, fast := pre, pre
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	l1Head := pre.Next
	l2Head := slow.Next
	slow.Next = nil
	l1 := sortList(l1Head)
	l2 := sortList(l2Head)
	newHead := merge(l1, l2)
	return newHead
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{}
	mark := head

	for ; l1 != nil && l2 != nil; mark = mark.Next {
		if l1.Val <= l2.Val {
			mark.Next = l1
			l1 = l1.Next
		} else {
			mark.Next = l2
			l2 = l2.Next
		}
	}
	if l1 == nil {
		mark.Next = l2
	} else {
		mark.Next = l1
	}
	return head.Next
}
