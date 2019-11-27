package kClosest

func kClosest(points [][]int, K int) [][]int {
	n, l, r := len(points), 0, len(points)-1
	if n == 0 {
		return [][]int{}
	}

	// divide
	mid := 0
	for l <= r {
		mid = helper(points, l, r) // divide, use pivot in quick sort
		if mid == K {
			break
		} else if mid < K {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return points[:K]
}

func helper(arr [][]int, start int, end int) int {
	i, j := start, start
	for ; j < end; j++ {
		if compare(arr[j], arr[end]) <= 0 {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}

func compare(l []int, r []int) int {
	return (l[0]*l[0] + l[1]*l[1]) - (r[0]*r[0] + r[1]*r[1])
}
