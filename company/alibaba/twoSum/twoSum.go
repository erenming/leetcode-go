package twoSum

func twoSum(nums []int, target int) []int {
	ht := make(map[int]int)
	for i, v := range nums {
		idx, ok := ht[v]
		if ok {
			return []int{idx, i}
		} else {
			ht[target-v] = i
		}
	}
	return nil
}
