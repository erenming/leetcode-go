package generateParenthesis

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	dp(&res, n, n, "")
	return res
}

func dp(res *[]string, left, right int, s string) {
	if left == 0 && right == 0 {
		*res = append(*res, s)
		return
	}
	if left > 0 {
		dp(res, left-1, right, s+"(")
		return
	}
	if right > left {
		dp(res, left, right-1, s+")")
		return
	}
}
