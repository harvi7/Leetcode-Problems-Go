package bfs

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}

	m, n := len(grid), len(grid[0])
	q := []Point{{0, 0, 1}}
	grid[0][0] = 1

	dirs := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	for len(q) > 0 {
		point := q[0]
		q = q[1:]

		if point.r == m-1 && point.c == n-1 {
			return point.dist
		}

		for _, d := range dirs {
			r, c := point.r+d[0], point.c+d[1]
			if r >= 0 && c >= 0 && r < m && c < n && grid[r][c] == 0 {
				q = append(q, Point{r, c, point.dist + 1})
				grid[r][c] = 1
			}
		}
	}

	return -1
}

type Point struct{ r, c, dist int }
