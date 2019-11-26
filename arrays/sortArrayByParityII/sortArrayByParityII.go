package sortArrayByParityII

func sortArrayByParityII(A []int) []int {
	i, j := 0, 1
	for i < len(A)-1 {
		if A[i]%2 == 1 { // odd
			A[i], A[j] = A[j], A[i]
			j += 2
		} else {
			i += 2
		}
	}
	return A
}
