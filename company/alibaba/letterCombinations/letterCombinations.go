package letterCombinations

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	s := newSolution()
	s.backtrack([]rune(digits), make([]rune, 0))
	return s.res
}

type solution struct {
	m   map[rune][]rune
	res []string
}

func newSolution() *solution {
	return &solution{
		m: map[rune][]rune{
			'2': {'a', 'b', 'c'},
			'3': {'d', 'e', 'f'},
			'4': {'g', 'h', 'i'},
			'5': {'j', 'k', 'l'},
			'6': {'m', 'n', 'o'},
			'7': {'p', 'q', 'r', 's'},
			'8': {'t', 'u', 'v'},
			'9': {'w', 'x', 'y', 'z'},
		},
	}
}

func (s *solution) backtrack(digits, r []rune) {
	if len(digits) <= 0 {
		s.res = append(s.res, string(r))
		return
	}
	dd := s.m[digits[0]]
	for _, num := range dd {
		s.backtrack(digits[1:], append(r, num))
	}
}
