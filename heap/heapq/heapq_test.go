package heapq_test

import (
	"leetcode-go/heap/heapq"
	"testing"
)

func TestHeap(t *testing.T) {
	h := heapq.New([]int{3, 2, 4, 1, 9})
	t.Log(h)
	h.Pop()
	t.Log(h)
	h.Push(5)
	t.Log(h)
	h.Pop()
	h.Pop()
	t.Log(h)
}
