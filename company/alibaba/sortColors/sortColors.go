package sortColors

func sortColors(nums []int) {
	n := len(nums)
	l, i, r := 0, 0, n-1
	for i <= r {
		if nums[i] == 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
			i++
		} else if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		}
	}
}
