package minPathSum

import (
	"math"
)
// version 2 recursion
func minPathSum(grid [][]int) int {

}


// func minPathSum(grid [][]int) int {
// 	m := len(grid)
// 	if m < 1 {
// 		return 0
// 	}
// 	n := len(grid[0])
// 	s := New(m, n, grid)
// 	return s.dp(0, 0)
// }
//
// type solution struct {
// 	m      int
// 	n      int
// 	matrix [][]int
// 	grid   [][]int
// }
//
// func New(row, col int, grid [][]int) *solution {
// 	m := make([][]int, row)
// 	for i := 0; i < row; i++ {
// 		m[i] = make([]int, col)
// 	}
// 	return &solution{row, col, m, grid}
// }
//
// func (s *solution) dp(i, j int) int {
// 	if i+1 == s.m && j+1 == s.n {
// 		return s.grid[i][j]
// 	}
// 	if s.matrix[i][j] != 0 {
// 		return s.matrix[i][j]
// 	}
// 	var res int
// 	if i+1 == s.m {
// 		res = s.dp(i, j+1)
// 	} else if j+1 == s.n {
// 		res = s.dp(i+1, j)
// 	} else {
// 		res = int(math.Min(float64(s.dp(i+1, j)), float64(s.dp(i, j+1))))
// 	}
// 	res += s.grid[i][j]
// 	s.matrix[i][j] = res
// 	return res
// }
