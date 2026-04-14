package top_k_frequent

import "container/heap"

type MinHeap [][2]int

func (m *MinHeap) Swap(i int, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MinHeap) Len() int {
	return len(*m)
}

func (m *MinHeap) Less(i int, j int) bool {
	return (*m)[i][1] < (*m)[j][1]
}

func (m *MinHeap) Push(i interface{}) {
	*m = append(*m, i.([2]int))
}

func (m *MinHeap) Pop() interface{} {
	old := *m
	last := old[len(old)-1]
	*m = old[:len(old)-1]
	return last

}

func topKFrequent(nums []int, k int) []int {
	freqs := make(map[int]int)

	for _, n := range nums {
		if freq, ok := freqs[n]; ok {
			freqs[n] = freq + 1
		} else {
			freqs[n] = 1
		}
	}

	m := &MinHeap{}

	for num, freq := range freqs {
		heap.Push(m, [2]int{num, freq})

		if m.Len() > k {
			heap.Pop(m)
		}
	}

	res := make([]int, 0, k)

	for _, pair := range *m {
		res = append(res, pair[0])
	}

	return res
}
