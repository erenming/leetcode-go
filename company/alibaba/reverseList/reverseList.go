package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := head
	pre.Next, head = nil, head.Next
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = pre
		pre = tmp
	}
	return pre
}
