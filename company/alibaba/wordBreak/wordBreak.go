package wordBreak

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]struct{}, len(wordDict))
	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]] = struct{}{}
	}
	memo := make(map[string]bool)
	return wb(s, wordMap, memo)
}

func wb(s string, wordMap map[string]struct{}, memo map[string]bool) bool {
	if v, ok := memo[s]; ok {
		return v
	}
	if _, ok := wordMap[s]; ok {
		memo[s] = true
		return true
	}
	for i := 1; i < len(s); i++ {
		left := s[:i]
		right := s[i:]
		if wb(left, wordMap, memo) && wb(right, wordMap, memo) {
			memo[s] = true
			return true
		}
	}
	memo[s] = false
	return false
}
