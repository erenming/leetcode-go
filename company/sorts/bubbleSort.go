package sorts

func bubbleSort(nums []int) []int {
	if nums == nil || len(nums) <= 1 {
		return nums
	}
	n := len(nums)
	for j := n - 1; j >= 0; j-- {
		for i := 0; i <= j-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	return nums
}
