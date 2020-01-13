package findKthLargest

func findKthLargest(nums []int, k int) int {
	var res int
	lo, hi := 0, len(nums)-1
	for {
		p := partition(nums, lo, hi)
		if p < k-1 {
			lo = p + 1
		} else if p > k-1 {
			hi = p - 1
		} else {
			res = nums[p]
			break
		}
	}
	return res
}

func partition(nums []int, lo, hi int) int {
	pivot, l, r := nums[lo], lo+1, hi
	for l <= r {
		if nums[l] < pivot && nums[r] > pivot {
			nums[l] ,nums[r] = nums[r], nums[l]
			l++
			r--
		}
		if nums[l] >= pivot {
			l++
		}
		if nums[r] <= pivot {
			r--
		}
	}
	nums[lo], nums[r] = nums[r], nums[lo]
	return r
}
