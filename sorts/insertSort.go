package sorts

func InsertSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 1; i < n-1; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}
