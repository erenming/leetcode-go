package romanToInt

func romanToInt(s string) int {
	if len(s) < 1 {
		return 0
	}
	ht := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	source := []rune(s)
	pre := source[0]
	res := ht[pre]
	for i := 1; i < len(source); i++ {
		c := source[i]
		res += ht[c]
		if ht[pre] < ht[c] {
			res -= 2 * ht[pre]
		}
		pre = c
	}
	return res
}
