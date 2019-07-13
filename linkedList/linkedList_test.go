package linkedList

import "testing"

func createLinkedList() *ListNode {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	return &head
}

func TestRemoveNthFromEnd(t *testing.T) {
	removeNthFromEnd(createLinkedList(), 2)
}
