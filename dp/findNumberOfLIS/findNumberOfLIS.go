package findNumberOfLIS

import (
	"math"
)

func findNumberOfLIS(nums []int) int {
	n, maxlen := len(nums), 0
	_cnt := make([]int, n)
	_len := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		maxlen = int(math.Max(float64(maxlen), float64(getLength(nums, i))))
	}


}

func getLength(nums []int, idx int) int {

}

func getCount(nums []int, idx int) int {

}
