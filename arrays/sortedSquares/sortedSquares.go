package sortedSquares

func sortedSquares(A []int) []int {
	n := len(A)
	res := make([]int, n)
	i, j, index := 0, n-1, n-1
	for i <= j {
		ii := A[i] * A[i]
		jj := A[j] * A[j]
		if ii > jj {
			res[index] = ii
			i++
		} else {
			res[index] = jj
			j--
		}
		index--
	}
	return res
}
