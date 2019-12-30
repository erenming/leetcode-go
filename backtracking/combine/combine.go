package combine

func combine(n int, k int) [][]int {
	s := New(n, k)
	s.dfs(0, []int{})
	return s.res
}

type solution struct {
	k   int
	arr []int
	res [][]int
}

func New(n, k int) *solution {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	return &solution{
		k:   k,
		arr: arr,
		res: make([][]int, 0),
	}
}

func (s *solution) dfs(start int, item []int) {
	for i := start; i < len(s.arr); i++ {
		newItem :=  append(item, s.arr[i])
		if len(newItem) == s.k {
			data := make([]int, len(newItem))
			copy(data, newItem)
			s.res = append(s.res, newItem)
			continue
		}
		s.dfs(i+1, newItem)
	}
}
