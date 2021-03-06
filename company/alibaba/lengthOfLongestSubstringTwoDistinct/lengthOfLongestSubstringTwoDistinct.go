package lengthOfLongestSubstringTwoDistinct

import (
	"math"
)

// eceba
func lengthOfLongestSubstringTwoDistinct(s string) int {
	source := []rune(s)
	n, i, j, res := len(source), 0, 0, 2
	if n < 3 {
		return n
	}
	hs := make(map[rune]int)
	for j < n {
		if len(hs) < 3 {
			hs[source[j]] = j
			j++
		}
		if len(hs) == 3 {
			leftMost := math.MaxInt64
			for _, v := range hs {
				if v < leftMost {
					leftMost = v
				}
			}
			delete(hs, source[leftMost])
			i = leftMost+1
		}
		res = int(math.Max(float64(j-i), float64(res)))
	}
	return res
}
