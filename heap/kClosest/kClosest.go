package main

import (
	"container/heap"
	"fmt"
)

// heap solution
type Point struct {
	x      int
	y      int
	square int
}

func NewPoint(x, y int) *Point {
	return &Point{
		x:      x,
		y:      y,
		square: x*x + y*y,
	}
}

type PointHeap []*Point

func (h PointHeap) Len() int {
	return len(h)
}

// max heap
func (h PointHeap) Less(i, j int) bool {
	return h[i].square > h[j].square
}

func (h PointHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PointHeap) Push(x interface{}) {
	*h = append(*h, x.(*Point))
}

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func kClosest(points [][]int, K int) [][]int {
	h := &PointHeap{}
	heap.Init(h)
	for i := 0; i < len(points); i++ {
		heap.Push(h, NewPoint(points[i][0], points[i][1]))
	}
	// 若nums元素个数大于k个,保留最小的k个元素在堆中
	for len(*h) > K {
		heap.Pop(h)
	}

	rv := make([][]int, K)
	for i := 0; i < K; i++ {
		p := heap.Pop(h).(*Point)
		rv[K-i-1] = []int{p.x, p.y}
	}
	return rv
}

func main() {
	a := [][]int{{3, 3}, {5, 1}, {-2, 4}}
	fmt.Println(kClosest(a, 2))

	b := [][]int{{1, 3}, {-2, 2}}
	fmt.Println(kClosest(b, 1))
}
