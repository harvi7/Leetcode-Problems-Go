package backtracking

func backtrack(grid [][]int, row, col, remain, rows, cols int) int {
	if grid[row][col] == 2 && remain == 1 {
		return 1
	}

	temp := grid[row][col]
	grid[row][col] = -4
	remain--

	rowOffsets := []int{0, 0, 1, -1}
	colOffsets := []int{1, -1, 0, 0}
	pathCount := 0

	for i := 0; i < 4; i++ {
		nextRow := row + rowOffsets[i]
		nextCol := col + colOffsets[i]

		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols {
			continue
		}
		if grid[nextRow][nextCol] < 0 {
			continue
		}
		pathCount += backtrack(grid, nextRow, nextCol, remain, rows, cols)
	}

	grid[row][col] = temp
	return pathCount
}

func uniquePathsIII(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	nonObstacles := 0
	startRow, startCol := 0, 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] >= 0 {
				nonObstacles++
			}
			if grid[r][c] == 1 {
				startRow, startCol = r, c
			}
		}
	}

	return backtrack(grid, startRow, startCol, nonObstacles, rows, cols)
}
