package removeDuplicates

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 1
	}
	pre, cur := 0, 1
	for cur < n {
		if nums[pre] == nums[cur] {
			cur++
		} else {
			pre++
			nums[pre] = nums[cur]
			cur++
		}
	}
	return pre+1
}
