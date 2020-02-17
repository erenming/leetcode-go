package removeDuplicates

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return len(nums)
	}
	n := len(nums)
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
