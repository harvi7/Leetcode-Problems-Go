func maxAreaOfIsland(grid [][]int) int {
    max_area := 0
    
    for r := 0; r < len(grid); r++ {
        for c := 0; c < len(grid[0]); c++ {
            if grid[r][c] == 1 {
                max_area = max(max_area, dfs(grid, r, c))
            }
        }
    }
    return max_area
}

func dfs(grid [][]int, r, c int) int {
    if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] == 1 {
        grid[r][c] = 0
        
        return 1 + dfs(grid, r - 1, c) + dfs(grid, r, c - 1) + dfs(grid, r + 1, c) + dfs(grid, r, c + 1)
    }
    
    return 0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}