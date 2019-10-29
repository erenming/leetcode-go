package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type Counter struct {
	char  string
	count int
}

func NewCounter(char string, count int) *Counter {
	return &Counter{char, count}
}

// max heap
type CounterHeap []*Counter

func (h CounterHeap) Len() int {
	return len(h)
}

func (h CounterHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h CounterHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *CounterHeap) Push(x interface{}) {
	*h = append(*h, x.(*Counter))
}

func (h *CounterHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func frequencySort(s string) string {
	table := make(map[string]int)
	for _, char := range strings.Split(s, "") {
		if _, ok := table[char]; ok {
			table[char]++
		} else {
			table[char] = 1
		}
	}
	h := &CounterHeap{}
	heap.Init(h)
	for key, value := range table {
		heap.Push(h, NewCounter(key, value))
	}
	var rv strings.Builder
	for h.Len() > 0 {
		e := heap.Pop(h).(*Counter)
		for i := 0; i < e.count; i++ {
			rv.WriteString(e.char)
		}
	}
	return rv.String()
}

func main() {
	fmt.Println(frequencySort("tree"))
}
