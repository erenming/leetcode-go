package sorts

func QuickSort(arr []int) {
	quickSork(arr, 0, len(arr)-1)
}

func quickSork(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(arr, start, end)
	quickSork(arr, start, pivot-1)
	quickSork(arr, pivot+1, end)
}

func partition(arr []int, start, end int) int {
	p := end
	i := start // < i的为小于等于arr[p]的数, [i, j)之间的为大于arr[p]的数
	j := start // >= j的为未筛选的数
	for ; j < p; j++ {
		if arr[j] <= arr[p] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
