package sorts

func QuickSort(nums []int) []int {
	if nums == nil || len(nums) <= 1 {
		return nums
	}
	n := len(nums)
	_sort(nums, 0, n-1)
	return nums
}

func _sort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := partition(nums, lo, hi)
	_sort(nums, lo, mid-1)
	_sort(nums, mid+1, hi)
}

func partition(nums []int, lo int, hi int) int {
	pivot := lo


}
