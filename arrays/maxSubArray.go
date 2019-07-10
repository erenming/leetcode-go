package arrays

import "math"

// TODO need understand
func maxSubArray(nums []int) int {
	maxEndingHere := nums[0]
	maxSoFar := nums[0]
	for i := 1; i < len(nums); i++ {
		maxEndingHere = int(math.Max(float64(maxEndingHere + nums[i]), float64(nums[i])))
		maxSoFar = int(math.Max(float64(maxEndingHere), float64(maxSoFar)))
	}
	return maxSoFar
}
