package maximalSquare

import (
	"math"
)

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	sums := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		sums[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sums[i][j] = int(matrix[i-1][j-1]-'0') + sums[i][j-1] + sums[i-1][j] - sums[i-1][j-1]
		}
	}

	res := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			k := int(math.Min(float64(m-i+1), float64(n-j+1)))
			for ; k > 0; k-- {
				sum := sums[i+k-1][j+k-1] - sums[i+k-1][j-1] - sums[i-1][j+k-1] + sums[i-1][j-1]
				if sum == k*k {
					if sum > res {
						res = sum
						break
					}
				}
			}
		}
	}
	return res
}
