package findKthLargest

func findKthLargest(nums []int, k int) int {
	n, l, r := len(nums), 0, len(nums)-1
	if n == 0 {
		return -1
	}

	mid := 0
	for l <= r {
		mid = partition(nums, l, r)
		if mid == k-1 {
			break
		}
		if mid < k -1 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return nums[mid]
}

func partition(nums []int, l int, r int) int {
	p, i, j := r, l, l
	for ; j < r; j++ {
		if nums[j] >= nums[p] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[p] = nums[p], nums[i]
	return i
}
