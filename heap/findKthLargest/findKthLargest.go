package findKthLargest

import "container/heap"

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

func findKthLargest(nums []int, k int) int {
	kl := Constructor(k, nums)
	return (*kl.nums)[0]
}
