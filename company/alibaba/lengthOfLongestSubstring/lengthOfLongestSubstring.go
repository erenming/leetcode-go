package lengthOfLongestSubstring

import (
	"math"
)

func lengthOfLongestSubstring(s string) int {
	source := []rune(s)
	hs := make(map[rune]struct{})
	res, i, j := 0, 0, 0
	for j < len(source) {
		if _, ok := hs[source[j]]; !ok {
			hs[source[j]] = struct{}{}
			j++
			res = int(math.Max(float64(j-i), float64(res)))
		} else {
			delete(hs, source[i])
			i++
		}
	}
	return res
}
