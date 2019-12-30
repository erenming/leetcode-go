package numTilePossibilities

func numTilePossibilities(tiles string) int {
	record := make([]int, 26)
	for _, c := range tiles { // 剪枝
		record[c-'A']++
	}
	return dfs(record)
}

func dfs(record []int) int {
	sum := 0
	for i := 0; i < 26; i++ {
		if record[i] == 0 { // 剪枝
			continue
		}
		sum++
		record[i]--
		sum += dfs(record)
		record[i]++ // 回溯清理
	}
	return sum
}
