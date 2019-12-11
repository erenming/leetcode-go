package letterCasePermutation

func letterCasePermutation(S string) []string {
	s := []rune(S)
	res := newCollection()
	backTrack(s, 0, res)
	return res.data
}

func backTrack(s []rune, index int, res *collection) {
	if index == len(s) {
		res.pushBack(string(s))
		return
	}
	// to most depth
	backTrack(s, index+1, res)

	// backtrack
	if isLower(s[index]) {
		s[index] = s[index] - 32
		backTrack(s, index+1, res)
	} else if isUpper(s[index]) {
		s[index] = s[index] + 32
		backTrack(s, index+1, res)
	}
}

type collection struct {
	data []string
}

func newCollection() *collection {
	return &collection{
		data: make([]string, 0),
	}
}

func (c *collection) pushBack(s string) {
	c.data = append(c.data, s)
}

func isLower(c rune) bool {
	if (c >= 'a' && c <= 'z') {
		return true
	} else {
		return false
	}
}

func isUpper(c rune) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	} else {
		return false
	}
}

