package matrixScore

func matrixScore(A [][]int) int {
	m := len(A)
	n := len(A[0])

	cols := make([]int, n)

	// property 1, make row maximize
	for i := 0; i < m; i++ {
		if A[i][0] == 0 {
			flipRow(A, i, n)
		}
		for j := 0; j < m; j++ {
			if A[i][j] == 1 {
				cols[j]++
			}
		}
	}

	// property 2, make col maximize
	for i := 0; i < n; i++ {
		if cols[i] <= m/2 {
			flipCol(A, i, m)
		}
	}

	// sum

}

func flipRow(m [][]int, row int, n int) {
	for i := 0; i < n; i++ {
		if m[row][i] == 0 {
			m[row][i] = 1
		} else {
			m[row][i] = 0
		}
	}
}

func flipCol(m [][]int, col int, n int) {

}
