package linkedList

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	// find middle
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != nil {
		slow = slow.Next
	}

	slow = reverseList(slow)
	fast = head
	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}

//func reverseList(head *ListNode) *ListNode {
//	var preHead *ListNode
//
//	for ; head != nil ; {
//		recordHead := head.Next
//		head.Next = preHead
//		preHead = head
//		head = recordHead
//	}
//	return preHead
//}
