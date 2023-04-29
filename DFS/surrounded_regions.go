func solve(board [][]byte)  {
    rows := len(board)
    cols := 0
    if rows > 0{
        cols = len(board[0])
    }

    for i := 0; i < cols; i++ {
        if board[0][i]=='O'{
            update(board, 0, i, rows, cols)
        }
        if board[rows - 1][i] == 'O'{
            update(board,rows - 1, i, rows, cols)
        }
    }

    for j := 0; j < rows; j++ {
        if board[j][0]=='O'{
            update(board, j, 0, rows, cols)
        }
        if board[j][cols - 1] == 'O' {
            update(board, j, cols - 1, rows, cols)
        }
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if board[i][j] == 'O'{
                board[i][j]='X';
            } else if board[i][j] == '1'{
                board[i][j]='O';
            }
        }
    }
}

func update(board [][]byte, i, j, rows, cols int) {
    if i < 0 || rows <= i || j < 0 || cols <= j {
        return
    }
    if  board[i][j] == 'O' {
        board[i][j] = '1'
        if i > 1 {
            update(board,i - 1,j, rows, cols)
        }
        if j > 1 {
            update(board,i,j - 1, rows, cols)
        }
        if i < rows - 1 {
            update(board,i + 1,j, rows, cols)
        }
        if i < cols - 1 {
            update(board, i, j + 1, rows, cols)
        }
    }
    
}