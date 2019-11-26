package isSubsequence

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	i, j := 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}

	return i == len(s) && j <= len(t)
}
