package minAddToMakeValid

func minAddToMakeValid(S string) int {
	left, right := 0, 0
	for _, c := range S {
		switch c {
		case '(':
			right++
		case ')':
			if right > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return left + right
}
