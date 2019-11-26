package twoCitySchedCost

import "sort"

func twoCitySchedCost(costs [][]int) int {
	diff := make(NodeSlice, len(costs))
	for i, v := range costs {
		diff[i] = &Node{v[1] - v[0], i}
	}
	sort.Sort(diff)

	res := 0
	for i := 0; i < len(diff); i++ {
		if i < len(diff) / 2 {
			res += costs[diff[i].index][0]
		} else {
			res += costs[diff[i].index][1]
		}
	}
	return res
}

type NodeSlice []*Node

func (p NodeSlice) Len() int           { return len(p) }
func (p NodeSlice) Less(i, j int) bool { return p[i].value > p[j].value }
func (p NodeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Node struct {
	value int
	index int
}
