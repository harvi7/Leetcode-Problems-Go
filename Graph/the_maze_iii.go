package graph

import "container/heap"

type State struct {
	row, col int
	dist     int
	path     string
}

type MinHeap []State

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].dist == h[j].dist {
		return h[i].path < h[j].path
	}
	return h[i].dist < h[j].dist
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(State))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

var directions = [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
var textDirections = []string{"l", "u", "r", "d"}

var m, n int

func findShortestWay(maze [][]int, ball []int, hole []int) string {
	m = len(maze)
	n = len(maze[0])

	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}

	hp := &MinHeap{}
	heap.Init(hp)
	heap.Push(hp, State{ball[0], ball[1], 0, ""})

	for hp.Len() > 0 {
		curr := heap.Pop(hp).(State)

		row, col := curr.row, curr.col

		if seen[row][col] {
			continue
		}
		if row == hole[0] && col == hole[1] {
			return curr.path
		}

		seen[row][col] = true

		for _, next := range getNeighbors(row, col, maze, hole) {
			heap.Push(hp, State{
				row:  next.row,
				col:  next.col,
				dist: curr.dist + next.dist,
				path: curr.path + next.path,
			})
		}
	}

	return "impossible"
}

func valid(row, col int, maze [][]int) bool {
	if row < 0 || row >= m || col < 0 || col >= n {
		return false
	}
	return maze[row][col] == 0
}

func getNeighbors(row, col int, maze [][]int, hole []int) []State {
	var neighbors []State

	for i := 0; i < 4; i++ {
		dy := directions[i][0]
		dx := directions[i][1]
		dir := textDirections[i]

		currRow, currCol := row, col
		dist := 0

		for valid(currRow+dy, currCol+dx, maze) {
			currRow += dy
			currCol += dx
			dist++
			if currRow == hole[0] && currCol == hole[1] {
				break
			}
		}

		neighbors = append(neighbors, State{
			row:  currRow,
			col:  currCol,
			dist: dist,
			path: dir,
		})
	}

	return neighbors
}
