package sorts

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, 0, mid)
	mergeSort(arr, mid+1, end)
	mergeArr(arr, start, mid, end)
}

func mergeArr(arr []int, start, mid, end int) {
	aux := make([]int, end-start+1)
	i := start
	j := mid + 1
	for k := 0; k < len(aux); k++ {
		if i > mid {
			aux[k] = arr[j]
			j++
		} else if j > end {
			aux[k] = arr[i]
			i++
		} else if arr[i] <= arr[j] {
			aux[k] = arr[i]
			i++
		} else {
			aux[k] = arr[j]
			j++
		}
	}
	copy(arr[start:end+1], aux)
}
