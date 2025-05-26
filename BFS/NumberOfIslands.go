func numIslands(grid [][]byte) int {
    count := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
          if grid[i][j] == '1' {
              count += 1
              callBFS(grid, i, j)
          }
        }
    }
    return count
}
    
func callBFS(grid [][]byte, i, j int) {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
        return
    }    
    grid[i][j] = '0'
    callBFS(grid, i + 1, j)
    callBFS(grid, i - 1, j)
    callBFS(grid, i, j + 1)
    callBFS(grid, i, j - 1)
}