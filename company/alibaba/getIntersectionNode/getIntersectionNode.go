package getIntersectionNode

type ListNode struct {
	Val  int
	Next *ListNode
}

// hash table implement
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	ht := make(map[*ListNode]bool)
// 	for headA != nil {
// 		ht[headA] = true
// 		headA = headA.Next
// 	}
// 	for headB != nil {
// 		if _, ok := ht[headB]; ok {
// 			break
// 		} else {
// 			headB = headB.Next
// 		}
// 	}
// 	return headB
// }

// two pointer
// len(A) + len(B) = len(B) + len(A)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	i, j := headA, headB
	for i != j {
		if i == nil {
			i = headB
		} else {
			i = i.Next
		}
		if j == nil {
			j = headA
		} else {
			j = j.Next
		}
	}
	return i
}
