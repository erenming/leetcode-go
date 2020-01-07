package findNumberOfLIS

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	_cnt, _len := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		_cnt[i] = 1
		_len[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if _len[j]+1 > _len[i] {
					_len[i] = _len[j] + 1
					_cnt[i] = _cnt[j]
				} else if _len[j]+1 == _len[i] {
					_cnt[i] += _cnt[j]
				}
			}
		}
	}

	maxLen := 0
	for _, v := range _len {
		if v > maxLen {
			maxLen = v
		}
	}

	res := 0
	for i, v := range _len {
		if v == maxLen {
			res += _cnt[i]
		}
	}

	return res
}
