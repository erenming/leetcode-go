package partitionLabels

func partitionLabels(S string) []int {
	runeS := []rune(S)
	res := make([]int, 0)
	hash := make([]int, 26)
	for i, c := range runeS {
		hash[c-'a'] = i
	}

	count := 0
	index, last := 0, hash[runeS[0]-'a']
	for index < len(runeS) {
		if index > last {
			res = append(res, count)
			count = 0
		}
		if hash[S[index]-'a'] > last {
			last = hash[runeS[index]-'a']
		}
		index++
		count++
	}
	res = append(res, count)
	return res
}
