package kth_largest_element_in_stream

import "container/heap"

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	k       int
	minHeap *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	minHeap := &IntHeap{}
	heap.Init(minHeap)
	kthLargest := KthLargest{k: k, minHeap: minHeap}

	for _, num := range nums {
		kthLargest.Add(num)
	}

	return kthLargest
}
func (kl *KthLargest) Add(val int) int {
	if kl.minHeap.Len() < kl.k {
		heap.Push(kl.minHeap, val)
	} else if val > (*kl.minHeap)[0] {
		heap.Pop(kl.minHeap)
		heap.Push(kl.minHeap, val)
	}

	return (*kl.minHeap)[0]
}
