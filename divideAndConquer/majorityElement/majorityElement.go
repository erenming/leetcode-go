package majorityElement

// divide and conquer
func majorityElement(nums []int) int {

	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		return nums[0]
	}
	// divide
	a := majorityElement(nums[:n/2])
	b := majorityElement(nums[n/2:])

	// conquer
	if a == b {
		return a
	}
	c := count(nums, a)
	if c > n/2 {
		return a
	} else {
		return b
	}
}

func count(ints []int, val int) int {
	count := 0
	for _, v := range ints {
		if v == val {
			count++
		}
	}
	return count
}

