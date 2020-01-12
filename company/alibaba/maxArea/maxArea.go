package maxArea

import (
	"math"
)

func maxArea(height []int) int {
	n := len(height)
	i, j := 0, n-1
	res := 0
	for i < j {
		res = int(math.Max(float64(res), math.Min(float64(height[i]), float64(height[j])) * float64(j - i)))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return res
}
