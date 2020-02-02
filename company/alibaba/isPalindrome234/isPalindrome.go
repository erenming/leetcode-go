package isPalindrome234

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		if nums[l] == nums[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}
