package permute

func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	visited := make([]bool, len(nums))
	dfs(&res, path, nums, visited)
	return res
}

func dfs(res *[][]int, path []int, nums []int, visited []bool) {
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i :=0; i<len(nums); i++ {
		if !visited[i] {
			visited[i] = true
			dfs(res, append(append([]int{}, path...), nums[i]), nums, visited)
			visited[i] = false
		}
	}
}