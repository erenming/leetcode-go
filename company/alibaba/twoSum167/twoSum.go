package twoSum167

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	if n < 1 {
		return nil
	}
	m := make(map[int]int)
	lo, hi := 0, n-1
	for {
		mid := (hi-lo)/2 + lo
		if numbers[mid] > target {
			hi = mid - 1
		} else {
			break
		}
	}
	for ; lo < n; lo++ {
		idx, ok := m[numbers[lo]]
		if ok {
			return []int{idx+1, lo+1}
		} else {
			m[target-numbers[lo]] = lo
		}
	}
	return nil
}
