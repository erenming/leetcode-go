package numIslands

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m < 1 {
		return 0
	}
	n := len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j, m, n, grid)
			}
		}
	}
	return res
}

func dfs(i, j, m, n int, grid [][]byte) {
	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(i+1, j, m, n, grid)
	dfs(i, j+1, m, n, grid)
	dfs(i-1, j, m, n, grid)
	dfs(i, j-1, m, n, grid)
}
