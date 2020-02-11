package sorts

func insertSort(nums []int) []int {
	if nums == nil || len(nums) <= 1 {
		return nums
	}
	n := len(nums)
	for j := 1; j < n; j++ {
		k := j
		for i := j-1; i >= 0; i-- {
			if nums[k] < nums[i] {
				nums[i], nums[k] = nums[k], nums[i]
				k = i
			} else {
				break
			}
		}
	}
	return nums
}
