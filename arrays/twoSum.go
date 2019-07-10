package arrays

// use a map
func twoSum(nums []int, target int) []int {
	dic := make(map[int]int)
	for i, e := range nums {
		val, ok := dic[e]
		if ok == true {
			return []int{val, i}
		} else {
			m := target - e
			dic[m] = i
		}
	}
	return []int{0, 0}
}
