package graph

import "container/heap"

type Edge struct {
	node int
	prob float64
}

type Item struct {
	node int
	prob float64
}

type MaxHeap []Item

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].prob > h[j].prob } // max-heap
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	adj := make([][]Edge, n)
	for i, e := range edges {
		src, dst := e[0], e[1]
		adj[src] = append(adj[src], Edge{dst, succProb[i]})
		adj[dst] = append(adj[dst], Edge{src, succProb[i]})
	}

	visited := make([]bool, n)
	pq := &MaxHeap{{node: start_node, prob: 1.0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		item := heap.Pop(pq).(Item)
		node, prob := item.node, item.prob

		if visited[node] {
			continue
		}
		visited[node] = true

		if node == end_node {
			return prob
		}

		for _, edge := range adj[node] {
			if !visited[edge.node] {
				heap.Push(pq, Item{node: edge.node, prob: prob * edge.prob})
			}
		}
	}

	return 0.0
}
