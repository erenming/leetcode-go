package binarySearch

// 查找最后一个小于等于给定值的元素
func FindLastLessEqual(value int, a []int) int {
	lo := 0
	hi := len(a) - 1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if value >= a[mid] {
			if (mid == len(a)-1) || (a[mid+1] > value) {
				return mid
			} else {
				lo = mid + 1
			}
		} else {
			hi = mid - 1
		}
	}
	return -1
}
