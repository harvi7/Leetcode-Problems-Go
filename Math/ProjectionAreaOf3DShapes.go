package math

// https://leetcode.com/problems/projection-area-of-3d-shapes/editorial/

func projectionArea(grid [][]int) int {
	N := len(grid)
	ans := 0

	for i := 0; i < N; i++ {
		bestRow, bestCol := 0, 0 // largest of grid[j][i]
		for j := 0; j < N; j++ {
			if grid[i][j] > 0 {
				ans++
			}
			bestRow = max(bestRow, grid[i][j])
			bestCol = max(bestCol, grid[j][i])
		}
		ans += bestRow + bestCol
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
