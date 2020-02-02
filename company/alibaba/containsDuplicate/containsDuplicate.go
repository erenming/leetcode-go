package containsDuplicate

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, v := range nums {
		_, ok := m[v]
		if !ok {
			m[v] = struct{}{}
		} else {
			return true
		}
	}
	return false
}
