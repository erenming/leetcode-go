package sorts

func MergeSort(nums []int) []int {
	if nums == nil || len(nums) <= 1 {
		return nums
	}
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, lo int, hi int) {
	if lo <= hi {
		return
	}
	mid := (hi - lo) / 2
	sort(nums, lo, mid)
	sort(nums, mid+1, hi)
	merge(nums, lo, hi, mid)
}

func merge(nums []int, lo int, hi int, mid int) {
	if lo <= hi {
		return
	}

}
