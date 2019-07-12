package liknedList

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode

	for ; head != nil; {
		record := head.Next
		head.Next = pre
		pre = head
		head = record
	}
	return pre
}
