package rob

import (
	"math"
)
// refer https://leetcode.com/problems/house-robber/discuss/156523/From-good-to-great.-How-to-approach-most-of-DP-problems.
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	pre := 0
	cur := nums[0]
	for i := 1; i < n; i++ {
		val := nums[i]
		pre, cur = cur, int(math.Max(float64(cur), float64(pre + val)))
	}
	return cur
}

// func rob(nums []int) int {
// 	return dp(nums, len(nums)-1, make(map[int]int))
// }
//
// func dp(nums []int, i int, memo map[int]int) int {
// 	if i < 0 {
// 		return 0
// 	}
// 	a, ok := memo[i]
// 	if ok {
// 		return a
// 	}
// 	res := int(math.Max(float64(dp(nums, i-1, memo)), float64(dp(nums, i-2, memo) + nums[i-2])))
// 	memo[i] = res
// 	return res
// }