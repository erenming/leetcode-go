package sorts

func SortSlice(nums []int, method string) []int {
	switch method {
	case "bubble":
		return bubbleSort(nums)
	case "select":
		return selectSort(nums)
	case "insert":
		return insertSort(nums)
	case "quick":
		return QuickSort(nums)
	}
	return nil
}

