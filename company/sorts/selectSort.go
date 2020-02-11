package sorts

func selectSort(nums []int) []int {
	if nums == nil || len(nums) <= 1 {
		return nums
	}
	n := len(nums)
	for i := 0; i < n; i++ {
		min := i
		for j := i+1; j < n; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}
