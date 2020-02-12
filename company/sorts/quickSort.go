package sorts

func quickSort(nums []int) []int {
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
	pivot := hi
	i, j := 0, hi-1
	for i <= j {
		if nums[i] > nums[pivot] {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	nums[i], nums[pivot] = nums[pivot], nums[i]
	return i
}
