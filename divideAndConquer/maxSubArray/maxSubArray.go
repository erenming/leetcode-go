package maxSubArray

import "math"

// divide and conquer
func maxSubArray(nums []int) int {
	return maxSub(nums, 0, len(nums)-1)
}

func maxSub(nums []int, left int, right int) int {
	if left > right {
		return int(math.MinInt64)
	}
	// divide
	mid, ml, mr := left+(right-left)/2, 0, 0
	lmax := maxSub(nums, left, mid-1)
	rmax := maxSub(nums, mid+1, right)

	// conquer
	sumLeft := 0
	for i := mid - 1; i >= left; i-- {
		sumLeft += nums[i]
		ml = int(math.Max(float64(sumLeft), float64(ml)))
	}
	sumRight := 0
	for i := mid + 1; i <= right; i++ {
		sumRight += nums[i]
		mr = int(math.Max(float64(sumRight), float64(mr)))
	}
	return int(math.Max(math.Max(float64(lmax), float64(rmax)), float64(ml+nums[mid]+mr)))
}
