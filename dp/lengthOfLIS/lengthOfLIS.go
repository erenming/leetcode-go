package lengthOfLIS

import (
	"math"
)

// iteration 24ms
func lengthOfLIS(nums []int) int {
	n, res := len(nums), 0
	record := make([]int, n) // 备忘录（用于剪枝）
	for j := 0; j < n; j++ {
		tmp := 0
		for i := 0; i < j; i++ {
			if nums[i] < nums[j] {
				tmp = int(math.Max(float64(tmp), float64(record[i])))
			}
		}
		record[j] = tmp + 1
	}
	for _, v := range record {
		if v > res {
			res = v
		}
	}
	return res
}

// // recursion 44ms
// func lengthOfLIS(nums []int) int {
// 	n, res := len(nums), 0
// 	record := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		res = int(math.Max(float64(res), float64(getLength(nums, i, &record))))
// 	}
// 	return res
// }
//
// func getLength(nums []int, r int, record *[]int) int {
// 	if r == 0 {
// 		return 1
// 	}
// 	if (*record)[r] > 0 {
// 		return (*record)[r]
// 	}
// 	res := 1
// 	for i := 0; i < r; i++ {
// 		if nums[i] < nums[r] {
// 			res = int(math.Max(float64(res), float64(getLength(nums, i, record))+1))
// 		}
// 	}
// 	(*record)[r] = res
// 	return res
// }
