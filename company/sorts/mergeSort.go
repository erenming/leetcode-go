package sorts

func mergeSort(nums []int) []int {
	if nums == nil || len(nums) <= 1 {
		return nums
	}
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	mid := (lo + hi) / 2
	sort(nums, lo, mid)
	sort(nums, mid+1, hi)
	merge(nums, lo, hi, mid)
}

func merge(nums []int, lo int, hi int, mid int) {
	tmp := make([]int, hi-lo+1)
	l, r := lo, mid+1
	for i := 0; i < len(tmp); i++ {
		if r > hi {
			tmp[i] = nums[l]
			l++
		} else if l > mid {
			tmp[i] = nums[r]
			r++
		} else if nums[l] < nums[r] {
			tmp[i] = nums[l]
			l++
		} else if nums[l] >= nums[r] {
			tmp[i] = nums[r]
			r++
		}
	}
	for i := 0; i < len(tmp); i++ {
		nums[lo] = tmp[i]
		lo++
	}
}
