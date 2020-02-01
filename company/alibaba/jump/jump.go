package jump

import (
	"math"
)

func jump(nums []int) int {
	jumps := 0
	curEnd, curFarthest := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		curFarthest = int(math.Max(float64(curFarthest), float64(i+nums[i])))
		if i == curEnd {
			jumps++
			curEnd = curFarthest
		}
	}
	return jumps
}
