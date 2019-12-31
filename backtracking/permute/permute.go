package permute

func permute(nums []int) [][]int {
	return dfs([]int{}, nums, len(nums))
}

func dfs(perm []int, remain []int, k int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(remain); i++ {
		newPerm := append(perm, remain[i])
		if len(newPerm) == k {
			data := make([]int, len(newPerm))
			copy(data, newPerm)
			res = append(res, data)
			continue
		}
		newRemain := append(append([]int{}, remain[:i]...), remain[i+1:]...)
		res = append(res, dfs(newPerm, newRemain, k)...)
	}
	return res
}
