package reconstructQueue

import "sort"

func reconstructQueue(people [][]int) [][]int {

	// property 1: h maximum and k min
	var queue mySlice = people
	sort.Sort(queue)

	// property 2: satisfy k
	for i := 0; i < len(people); i++ {
		tmp := queue[i]
		index := queue[i][1]
		copy(queue[index+1:i+1], queue[index:i])
		queue[index] = tmp
	}

	return queue
}

type mySlice [][]int

func (p mySlice) Len() int { return len(p) }
func (p mySlice) Less(i, j int) bool {
	if p[i][0] > p[j][0] {
		return true
	} else if p[i][0] == p[j][0] {
		if p[i][1] <= p[j][1] {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
func (p mySlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
