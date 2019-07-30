package binarySearch

func searchRange(nums []int, target int) []int {
	lo, hi := 0, len(nums)-1

	boundary := []int{-1, -1}
	// search left boundary
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if target > nums[mid] {
			lo = mid + 1
		} else if target < nums[mid] {
			hi = mid - 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				boundary[0] = mid
				break
			} else {
				hi = mid - 1
			}
		}
	}

	lo, hi = 0, len(nums)-1
	// search right boundary
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if target > nums[mid] {
			lo = mid + 1
		} else if target < nums[mid] {
			hi = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				boundary[1] = mid
				break
			} else {
				lo = mid + 1
			}
		}

	}
	return boundary
}
