package sorts

func SortSlice(nums []int, method string) []int {
	switch method {
	case "bubble":
		return bubbleSort(nums)
	case "select":
		return selectSort(nums)
	case "insert":
		return insertSort(nums)
	case "merge":
		return mergeSort(nums)
	case "quick":
		return quickSort(nums)
	}
	return nil
}
