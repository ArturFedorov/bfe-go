package find_median_from_data_stream

import "container/heap"

type MaxHeap []int

func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

type MinHeap []int

func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

type MedianFinder struct {
	low  *MaxHeap
	high *MinHeap
}

func Constructor() MedianFinder {
	l, h := &MaxHeap{}, &MinHeap{}
	heap.Init(l)
	heap.Init(h)
	return MedianFinder{l, h}
}

func (mf *MedianFinder) AddNum(num int) {
	heap.Push(mf.low, num)
	heap.Push(mf.high, heap.Pop(mf.low))

	if mf.low.Len() < mf.high.Len() {
		heap.Push(mf.low, heap.Pop(mf.high))
	}
}
func (mf *MedianFinder) FindMedian() float64 {
	if mf.low.Len() > mf.high.Len() {
		return float64((*mf.low)[0])
	}

	return (float64((*mf.low)[0]) + float64((*mf.high)[0])) / 2.0
}
