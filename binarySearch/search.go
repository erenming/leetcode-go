package binarySearch

func search(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if nums[mid] == target {
			return mid
		}
		if nums[lo] <= nums[mid] {
			if target >= nums[lo] && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else if nums[mid] <= nums[hi] {
			if target <= nums[hi] && target > nums[mid] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
