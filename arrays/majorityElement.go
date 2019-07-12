package arrays

func majorityElement(nums []int) int {
	var count, major int
	count = 0
	for _, e := range nums {
		if count == 0 {
			major = e
			count++
			continue
		}

		if e == major {
			count++
		} else {
			count--
		}
	}
	return major
}
