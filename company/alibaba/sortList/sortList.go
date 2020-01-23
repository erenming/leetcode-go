package sortList

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	r := slow.Next
	l := head
	slow.Next = nil
	left := sortList(l)
	right := sortList(r)
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	list := &ListNode{-1, nil}
	head := list
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			list.Next, l1 = l1, l1.Next
		} else {
			list.Next, l2 = l2, l2.Next
		}
		list = list.Next
	}
	if l1 == nil {
		list.Next = l2
	}
	if l2 == nil {
		list.Next = l1
	}
	return head.Next
}