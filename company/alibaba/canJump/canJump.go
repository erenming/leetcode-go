package canJump

import (
	"math"
)

func canJump(nums []int) bool {
	n := len(nums)
	reach, i := 0, 0
	for ; i < n && i <= reach; i++ {
		reach = int(math.Max(float64(reach), float64(nums[i]+i)))
	}
	return i == n
}
