package firstUniqChar

func firstUniqChar(s string) int {
	m := make(map[rune][2]int)
	for i, v := range []rune(s) {
		if val, ok := m[v]; ok {
			val[1]++
			m[v] = val
		} else {
			m[v] = [2]int{i, 1}
		}
	}
	first := -1
	for _, v := range m {
		if v[1] == 1 && (first == -1 || v[0] < first){
			first = v[0]
		}
	}
	return first
}
