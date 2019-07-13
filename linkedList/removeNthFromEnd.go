package linkedList

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentry := &ListNode{0, nil}
	sentry.Next = head
	start := sentry
	slow := start
	fast := start
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	for ; fast != nil; {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next  // delete node
	return start.Next
}
