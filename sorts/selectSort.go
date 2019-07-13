package sorts

func SelectSort(a []int, n int) {
	for i := n - 1; i > 1; i-- {
		max_i := 0
		for j := 1; j < i+1; j++ {
			if a[max_i] < a[j] {
				max_i = j
			}
		}
		a[max_i], a[i] = a[i], a[max_i]
	}
}
