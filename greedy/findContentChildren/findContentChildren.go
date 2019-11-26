package findContentChildren

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] { // content
			res++
			i++
		}
		j++
	}
	return res
}
