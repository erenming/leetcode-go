package generateParenthesis

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	dfs(n, n, "", &res)
	return res
}

func dfs(left, right int, str string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, str)
		return
	}
	if left > 0 {
		dfs(left-1, right, str+"(", res)
	}
	if right > left { // 剪枝逻辑
		dfs(left, right-1, str+")", res)
	}
}
