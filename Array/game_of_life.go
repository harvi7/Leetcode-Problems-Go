package array

var dirs = [][]int{
	{0, -1}, {0, 1}, {1, 0}, {-1, 0},
	{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
}

func gameOfLife(board [][]int) {
	if board == nil || len(board) == 0 {
		return
	}
	rows, cols := len(board), len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 0 {
				lives := count(board, i, j)
				if lives == 3 {
					board[i][j] = 2 // dead -> live
				}
			} else if board[i][j] == 1 {
				lives := count(board, i, j)
				if lives < 2 || lives > 3 {
					board[i][j] = -1 // live -> die
				}
			}
		}
	}
	update(board)
}

func update(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 1 // new live
			} else if board[i][j] == -1 {
				board[i][j] = 0 // new dead
			}
		}
	}
}

func count(board [][]int, row, col int) int {
	res := 0
	for _, dir := range dirs {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if newRow >= 0 && newRow < len(board) &&
			newCol >= 0 && newCol < len(board[0]) &&
			(board[newRow][newCol] == 1 || board[newRow][newCol] == -1) {
			res++
		}
	}
	return res
}
