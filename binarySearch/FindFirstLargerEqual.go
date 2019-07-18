package binarySearch

// 查找第一个大于等于给定值的元素
func FindFirstLargerEqual(value int, a []int) int {
	lo := 0
	hi := len(a) - 1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if value <= a[mid] {
			if (mid == 0) || (a[mid-1] < value) {
				return mid
			} else {
				hi = mid-1
			}
		} else {
			lo = mid+1
		}
	}
	return -1
}
