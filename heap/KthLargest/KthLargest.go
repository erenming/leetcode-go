package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

type KthLargest struct {
	nums *IntHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
	}
	// 若nums元素个数大于k个,保留最大的k个元素
	for len(*h) > k {
		heap.Pop(h)
	}
	return KthLargest{
		nums: h,
		k:    k,
	}
}

func (this *KthLargest) Add(val int) int {
	if len(*this.nums) < this.k {
		heap.Push(this.nums, val)
	} else if val > (*this.nums)[0] {
		heap.Pop(this.nums)
		heap.Push(this.nums, val)
	}
	return (*this.nums)[0]
}

func main() {
	obj := Constructor(3, []int{1, 7, 5, 2, 4});
	param_1 := obj.Add(10);
	fmt.Println(param_1)
	fmt.Println(obj)
}
/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
