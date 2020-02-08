package maxAreaOfIsland

import (
	"math"
)

func maxAreaOfIsland(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	maxArea := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			area := dfs(grid, i, j, m, n, 0)
			maxArea = int(math.Max(float64(maxArea), float64(area)))
		}
	}
	return maxArea
}

func dfs(grid [][]int, i int, j int, m int, n int, area int) int {
	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
		return area
	}
	grid[i][j] = 0
	area++
	area = dfs(grid, i-1, j, m, n, area)
	area = dfs(grid, i, j-1, m, n, area)
	area = dfs(grid, i+1, j, m, n, area)
	area = dfs(grid, i, j+1, m, n, area)
	return area
}
