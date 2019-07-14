package sorts

func FindKth(arr []int, k int) int {
	return findKth(arr, k, 0, len(arr)-1)

}

func findKth(arr []int, k, start, end int) int {
	pivot := rpartition(arr, start, end)
	if pivot+1 == k {
		return arr[pivot]
	} else if pivot+1 > k { // 左边
		return findKth(arr, k, start, pivot-1)
	} else {
		return findKth(arr, k, pivot+1, end)
	}
}

func rpartition(arr []int, start, end int) int {
	p := end
	i := start // < i的为大于等于arr[p]的数, [i, j)之间的为小于arr[p]的数
	j := start // >= j的为未筛选的数
	for ; j < p; j++ {
		if arr[j] >= arr[p] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
