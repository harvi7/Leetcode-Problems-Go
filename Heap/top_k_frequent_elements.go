package heap

import "container/heap"

type Pair struct {
	num  int
	freq int
}

// MaxHeap implements heap.Interface for []Pair (max-heap by freq)
type MaxHeapTopK []Pair

func (h MaxHeapTopK) Len() int           { return len(h) }
func (h MaxHeapTopK) Less(i, j int) bool { return h[i].freq > h[j].freq } // max-heap
func (h MaxHeapTopK) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeapTopK) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MaxHeapTopK) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	if k == len(nums) {
		return nums
	}

	// Count frequencies
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	// Build max heap
	h := &MaxHeapTopK{}
	for num, freq := range count {
		heap.Push(h, Pair{num, freq})
	}

	// Extract top k elements
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(h).(Pair).num
	}

	return res
}
