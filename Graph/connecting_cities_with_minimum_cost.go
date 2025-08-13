package graph

import "sort"

type DisjointSet struct {
	weights []int
	parents []int
}

func NewDisjointSet(N int) *DisjointSet {
	weights := make([]int, N+1)
	parents := make([]int, N+1)
	for i := 1; i <= N; i++ {
		parents[i] = i
		weights[i] = 1
	}
	return &DisjointSet{
		weights: weights,
		parents: parents,
	}
}

func (ds *DisjointSet) Find(a int) int {
	for a != ds.parents[a] {
		// Path compression
		ds.parents[a] = ds.parents[ds.parents[a]]
		a = ds.parents[a]
	}
	return a
}

func (ds *DisjointSet) Union(a, b int) {
	rootA := ds.Find(a)
	rootB := ds.Find(b)
	if rootA == rootB {
		return
	}

	// Weighted union
	if ds.weights[rootA] > ds.weights[rootB] {
		ds.parents[rootB] = rootA
		ds.weights[rootA] += ds.weights[rootB]
	} else {
		ds.parents[rootA] = rootB
		ds.weights[rootB] += ds.weights[rootA]
	}
}

func (ds *DisjointSet) isInSameGroup(a, b int) bool {
	return ds.Find(a) == ds.Find(b)
}

func minimumCost(n int, connections [][]int) int {
	ds := NewDisjointSet(n)

	// Sort connections by weight
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	total := 0
	cost := 0

	for _, conn := range connections {
		a, b := conn[0], conn[1]
		if ds.isInSameGroup(a, b) {
			continue
		}
		ds.Union(a, b)
		cost += conn[2]
		total++
	}

	if total == n-1 {
		return cost
	}
	return -1
}
