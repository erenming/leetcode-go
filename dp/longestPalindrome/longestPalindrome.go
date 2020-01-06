package longestPalindrome

func longestPalindrome(s string) string {
	chars := []rune(s)
	n := len(chars)
	res := ""

	dp := make([][]bool, 0)
	for i := 0; i < n; i++ {
		dp = append(dp, make([]bool, n))
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			tmp := (chars[i] == chars[j]) && (j-i < 3 || dp[i+1][j-1])
			dp[i][j] = tmp
			if dp[i][j] && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
