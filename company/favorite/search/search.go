package search

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	n := len(nums)
	lo, hi := 0, n-1
	for lo < hi {
		mid := (hi + lo) / 2
		if target == nums[mid] {
			return mid
		}
		if nums[lo] <= nums[mid] { // left side ordered
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else { // right side ordered
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	if nums[lo] == target {
		return lo
	} else {
		return -1
	}
}
