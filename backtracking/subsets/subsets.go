package subsets

func subsets(nums []int) [][]int {
	res := &[][]int{{}}
	dfs(res, nums, []int{})
	return *res
}

func dfs(res *[][]int, remain []int, cur []int) {
	if len(remain) == 0 {
		return
	}
	for i, val := range remain {
		tmp :=  newSlice(cur, val)
		*res = append(*res, tmp)
		dfs(res, append([]int{}, remain[i+1:]...), tmp)
	}
}

func newSlice(s []int, appendTo int) []int {
	res := make([]int, len(s)+1)
	copy(res, s)
	res[len(s)] = appendTo
	return res
}
