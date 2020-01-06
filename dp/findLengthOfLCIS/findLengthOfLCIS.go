package findLengthOfLCIS

import (
	"math"
)

func findLengthOfLCIS(nums []int) int {
	cnt, res := 0, 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > nums[i-1] {
			cnt++
			res = int(math.Max(float64(res), float64(cnt)))
		} else {
			cnt = 1
		}
	}
	return res
}
