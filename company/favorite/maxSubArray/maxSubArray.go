package maxSubArray

import (
	"math"
)

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	n := len(nums)
	max := nums[0]
	sum := nums[0]
	for i := 1; i < n; i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		max = int(math.Max(float64(max), float64(sum)))
	}
	return max
}
