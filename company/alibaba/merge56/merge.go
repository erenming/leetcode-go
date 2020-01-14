package merge56

import (
	"math"
	"sort"
)

// sort then merge
func merge(intervals [][]int) [][]int {
	sort.Sort(block(intervals))
	n := len(intervals)
	if n <= 1 {
		return intervals
	}
	start, end := intervals[0][0], intervals[0][1]
	res := make([][]int, 0)
	for i := 1; i < n; i++ {
		if end < intervals[i][0] {
			res = append(res, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		} else {
			end = int(math.Max(float64(end), float64(intervals[i][1])))
		}
	}
	res = append(res, []int{start, end})
	return res
}

type block [][]int

func (b block) Len() int {
	return len(b)
}

func (b block) Less(i, j int) bool {
	return b[i][0] < b[j][0]
}

func (b block) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
