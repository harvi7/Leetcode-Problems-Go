package heap

import (
	"container/heap"
	"sort"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)

	// Combine into pairs [capital, profit]
	projects := make([][2]int, n)
	for i := 0; i < n; i++ {
		projects[i] = [2]int{capital[i], profits[i]}
	}

	// Sort by capital
	sort.Slice(projects, func(i, j int) bool {
		return projects[i][0] < projects[j][0]
	})

	h := &MaxHeap{}
	heap.Init(h)

	ptr := 0
	for i := 0; i < k; i++ {
		// Add all affordable projects to max heap
		for ptr < n && projects[ptr][0] <= w {
			heap.Push(h, projects[ptr][1])
			ptr++
		}

		if h.Len() == 0 {
			break
		}

		// Pick most profitable project
		w += heap.Pop(h).(int)
	}

	return w
}
