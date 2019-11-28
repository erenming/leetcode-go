package diffWaysToCompute

import (
	"strconv"
)

func diffWaysToCompute(input string) []int {
	res := make([]int, 0)
	source := []rune(input)
	for i := 0; i < len(source); i++ {
		cur := source[i]
		if cur == '+' || cur == '-' || cur == '*' {
			p1 := diffWaysToCompute(string(source[:i]))
			p2 := diffWaysToCompute(string(source[i+1:]))
			for _, l := range p1 {
				for _, r := range p2 {
					if cur == '+' {
						res = append(res, l+r)
					} else if cur == '-' {
						res = append(res, l-r)
					} else {
						res = append(res, l*r)
					}
				}
			}
		}
	}
	if len(res) == 0 { // string only have number
		n, err := strconv.Atoi(input)
		if err != nil {
			return []int{}
		}
		res = append(res, n)
	}
	return res
}
