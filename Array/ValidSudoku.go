func isValidSudoku(board [][]byte) bool {
    N := 9

    row := [9][9]int{}       
    col := [9][9]int{}    
    boxes := [9][9]int{} 

    for r := 0; r < N; r++ {
        for c := 0; c < N; c++ {
            if (board[r][c] == '.') {
                continue;
            }
            pos := board[r][c] - '1'

            if row[r][pos] == 1 {
                return false
            }
            row[r][pos] = 1

            if col[c][pos] == 1 {
                return false
            }
            col[c][pos] = 1

            idx := (r / 3) * 3 + c / 3
            if boxes[idx][pos] == 1 {
                return false
            }
            boxes[idx][pos] = 1
        }
    }
    return true
}