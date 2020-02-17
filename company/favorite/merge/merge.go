package merge

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Sort(block(intervals))
	if len(intervals) < 1 {
		return intervals
	}
	n := len(intervals)
	res := [][]int{intervals[0]}
	for i := 1; i < n; i++ {
		last := res[len(res)-1]
		if last[1] >= intervals[i][0] {
			end := int(math.Max(float64(last[1]), float64(intervals[i][1])))
			res[len(res)-1] = []int{last[0], end}
		} else {
			res = append(res, intervals[i])
		}
	}
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
