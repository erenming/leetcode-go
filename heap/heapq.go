package heap

import (
	"fmt"
	"strconv"
	"strings"
)

// big heap
type Heap struct {
	array []int
}

func New(a []int) *Heap {
	h := &Heap{a}
	n := len(a)
	for i := n/2 - 1; i >= 0; i-- {
		h.bubbleDown(i)
	}
	return h
}

func (h *Heap) bubbleDown(i int) {
	n := len(h.array)
	for {
		j1 := 2*i + 1 // j1 is left
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		// j2 is right
		if j2 := j1 + 1 ; j2 < n && h.array[j1] < h.array[j2] {
			j = j2
		}
		if h.array[i] < h.array[j] {
			h.swap(i, j)
		}
		i = j
	}
}

func (h *Heap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *Heap) bubbleUp(i int) {
	for {
		j := (i - 1) / 2 // parent
		if i == j || h.array[j] > h.array[i] {
			break
		}
		h.swap(i, j)
		i = j
	}
}

func (h *Heap) Empty() bool {
	return len(h.array) == 0
}

func (h *Heap) Pop() int {
	var res int
	n := len(h.array)
	res, h.array[0] = h.array[0], h.array[n-1]

	h.array = h.array[:n-1]
	h.bubbleDown(0)
	return res
}

func (h *Heap) Push(item int) {
	h.array = append(h.array, item)
	h.bubbleUp(len(h.array) - 1)
}

func (h *Heap) String() string {
	var s strings.Builder
	for i := 0; i < len(h.array) - 1; i++ {
		s.WriteString(fmt.Sprintf("%d-", h.array[i]))
	}
	s.WriteString(strconv.Itoa(h.array[len(h.array) - 1]))
	return s.String()
}
