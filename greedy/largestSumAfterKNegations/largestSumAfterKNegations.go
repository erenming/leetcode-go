package largestSumAfterKNegations

import "sort"

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	index := 0
	for K > 0 {
		A[index] = -A[index]
		K--
		if index + 1 < len(A) && A[index+1] < A[index] {
			index++
		}
	}
	return sum(A)
}

func sum(s []int) (res int) {
	for _, v := range s {
		res += v
	}
	return res
}
