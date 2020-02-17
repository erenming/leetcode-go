package canJump

import (
	"math"
)

func canJump(nums []int) bool {
	n := len(nums)
	i, reach := 0, 0
	for ; i < n && i <= reach; i++ {
		reach = int(math.Max(float64(reach), float64(i + nums[i])))
	}
	return i == n
}
