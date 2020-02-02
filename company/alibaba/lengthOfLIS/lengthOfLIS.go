package lengthOfLIS

import (
	"math"
)

func lengthOfLIS(nums []int) int {
	n, res := len(nums), 0
	record := make([]int, n)
	for i := 0; i < n; i++ {
		tmp := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				tmp = int(math.Max(float64(tmp), float64(record[j]+1)))
			}
		}
		record[i] = tmp
	}
	for _, v := range record {
		if v > res {
			res = v
		}
	}
	return res
}
