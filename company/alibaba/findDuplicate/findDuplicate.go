package findDuplicate

func findDuplicate(nums []int) int {
	n := len(nums)
	for i := 0; i < n; {
		v := nums[i]
		if v != i {
			if nums[v] == v {
				return v
			} else {
				nums[i], nums[v] = nums[v], nums[i]
			}
		} else {
			i++
		}
	}
	return -1
}
