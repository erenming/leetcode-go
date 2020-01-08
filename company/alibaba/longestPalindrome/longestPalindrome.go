package longestPalindrome

import (
	"math"
)

func longestPalindrome(s string) string {
	source := []rune(s)
	n := len(source)
	if n < 1 {
		return s
	}
	_len := 0
	start, end := 0, 0
	for i := 0; i < n; i++ {
		odd := expendMiddle(source, i, i)
		even := expendMiddle(source, i, i+1)
		_len = int(math.Max(float64(odd), float64(even)))
		if _len > end-start+1 {
			start = i - (_len-1)/2
			end = i + _len/2
		}
	}
	return string(source[start:end+1])
}

func expendMiddle(s []rune, i, j int) int {
	if len(s) == 0 || i > j {
		return 0
	}
	for i >= 0 && j <= len(s)-1 && s[i] == s[j] {
		i--
		j++
	}
	return j - i -1
}
