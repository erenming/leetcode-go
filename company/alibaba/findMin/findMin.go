package findMin

import (
	"math"
)

func findMin(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	min := math.MaxInt64
	lo, hi := 0, n-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] <= nums[hi] {
			min = int(math.Min(float64(min), float64(nums[mid])))
			hi = mid-1
		} else {
			lo = mid+1
		}
	}
	return min

}
