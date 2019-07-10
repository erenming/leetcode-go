package arrays

import "math"

// relation between value in [1, n] and index in [0, n-1]
func findDisappearedNumbers(nums []int) []int {
	rv := []int{}
	for _, i := range nums {
		val := int(math.Abs(float64(i))) - 1
		if nums[val] > 0 {
			nums[val] = -nums[val]
		}
	}

	for index, j := range nums {
		if j > 0 {
			rv = append(rv, index+1)
		}
	}

	return rv
}
