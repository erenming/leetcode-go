package minDistance

import (
	"math"
)

func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	s := New(l1+1, l2+1)
	return s.dp(word1, word2, l1, l2)
}

type solution struct {
	matrix [][]int
}

func New(row, col int) *solution {
	m := make([][]int, row)
	for i := 0; i < row; i++ {
		tmp := make([]int, col)
		for j := 0; j < col; j++ {
			tmp[j] = -1
		}
		m[i] = tmp
	}
	return &solution{m}
}

func (s *solution) dp(word1, word2 string, l1, l2 int) int {
	if l1 == 0 {
		return l2
	}
	if l2 == 0 {
		return l1
	}
	if s.matrix[l1][l2] != -1 {
		return s.matrix[l1][l2]
	}
	res := -1
	if word1[l1-1] == word2[l2-1] {
		res = s.dp(word1, word2, l1-1, l2-1)
	} else {
		res = 1 + int(math.Min(float64(s.dp(word1, word2, l1-1, l2)), math.Min(float64(s.dp(word1, word2, l1, l2-1)), float64(s.dp(word1, word2, l1-1, l2-1)))))
	}
	s.matrix[l1][l2] = res
	return res
}
