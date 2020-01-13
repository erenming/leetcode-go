package search

func search(nums []int, target int) int {
	n := len(nums)
	if n < 1 {
		return -1
	}
	lo, hi := 0, n-1
	for lo < hi {
		mid := (hi + lo) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[lo] <= nums[mid] { // 左侧有序递增
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else { // 右侧有序
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	if target == nums[lo] {
		return lo
	} else {
		return -1
	}
}
