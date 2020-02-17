package jump

import (
	"math"
)

func jump(nums []int) int {
	res := 0
	n := len(nums)
	curEnd, curFarthest := 0, 0
	for i := 0; i < n-1; i++ {
		curFarthest = int(math.Max(float64(curFarthest), float64(i+nums[i])))
		if i == curEnd {
			res++
			curEnd = curFarthest
		}
	}
	return res
}
