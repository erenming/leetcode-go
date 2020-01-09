package threeSum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		target := -nums[i]
		left, right := i+1, n-1
		for left < right {
			if target < nums[left] + nums[right] {
				right--
			} else if target > nums[left] + nums[right] {
				left++
			} else {
				tmp := make([]int, 3)
				tmp[0] = nums[i]
				tmp[1] = nums[left]
				tmp[2] = nums[right]
				res = append(res, tmp)
				// process duplicates
				for left < right && nums[left] == tmp[1] {
					left++
				}
				for right > left && nums[right] == tmp[2] {
					right--
				}
			}
			for i < n-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res
}
