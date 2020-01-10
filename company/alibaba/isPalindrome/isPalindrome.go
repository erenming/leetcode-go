package isPalindrome

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xx := x
	res := 0
	for x != 0 {
		pop := x % 10
		x = x / 10
		res = res*10 + pop
	}
	return res == xx
}
