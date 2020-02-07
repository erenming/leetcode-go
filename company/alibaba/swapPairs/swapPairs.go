package swapPairs

type ListNode struct {
	Val  int
	Next *ListNode
}

	func swapPairs(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		pre, cur := head, head.Next
		for cur != nil && cur.Next != nil {
			pre.Val, cur.Val = cur.Val, pre.Val
			pre = pre.Next.Next
			cur = cur.Next.Next
		}
		if cur != nil {
			pre.Val, cur.Val = cur.Val, pre.Val
		}
		return head
	}
