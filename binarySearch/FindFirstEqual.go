package binarySearch

// 查找第一个等于给定值的元素的下标
func FindFirstEqual(value int, a []int) int {
	lo := 0
	hi := len(a) - 1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if value < a[mid] {
			hi = mid - 1
		} else if value > a[mid] {
			lo = mid + 1
		} else {
			if (mid == 0) || (a[mid-1] != value) {
				return mid
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
