package sorts

import "sort"

func Merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	interLen := len(intervals)

	sort.Sort(ByStart(intervals))

	for i := 0; i < interLen-1; {
		cur, next := intervals[i], intervals[i+1]
		if next[0] <= cur[1] {
			if next[1] > cur[1] {
				intervals[i][1] = intervals[i+1][1]
			}
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			interLen--
			continue
		}
		i++

	}
	return intervals
}

type ByStart [][]int

func (this ByStart) Len() int {
	return len(this)
}

func (this ByStart) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this ByStart) Less(i, j int) bool {
	return this[i][0] < this[j][0]
}
