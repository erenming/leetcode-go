package sorts

func getMax(a []int) int {
	max := a[0]
	for _, e := range a[1:] {
		if e > max {
			max = e
		}
	}
	return max
}

func BucketSort(a []int) {
	if len(a) <= 1 {
		return
	}
	max := getMax(a)
	k := max
	buckets := make([][]int, k)
	for _, e := range a {
		index := (e * (k - 1)) / max
		buckets[index] = append(buckets[index], e)
	}
	for i := range buckets {
		QuickSort(buckets[i])
	}
	pos := 0
	for i := range buckets {
		for _, e := range buckets[i] {
			a[pos] = e
			pos++
		}
	}

}
