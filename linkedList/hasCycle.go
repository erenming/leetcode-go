package linkedList


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func hasCycle(head *ListNode) bool {
	runner := head
	walker := head
	for ; runner != nil && runner.Next != nil; {
		runner = runner.Next.Next
		walker = walker.Next
		if walker == runner {
			return true
		}
	}
	return false
}