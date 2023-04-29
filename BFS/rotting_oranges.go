func orangesRotting(grid [][]int) int {
    rotten, r, c := [][2]int{}, len(grid), len(grid[0])
    for i := 0; i < r; i++ {
        for j := 0; j < c; j++ {
            if grid[i][j] == 2 { rotten = append(rotten, [2]int{i, j}) }
        }
    }
    timer := 0
    for {
        n := len(rotten)
        for i := 0; i < n; i++ {
            pop := rotten[0]
            rotten = rotten[1:]
            x, y := pop[0], pop[1]
            d := [][]int{
				[]int{x - 1, y},
				[]int{x + 1, y},
				[]int{x, y - 1},
				[]int{x, y + 1},
			}
            for _, v := range d {
                if v[0] < 0 || v[1] < 0 || v[0] >= r || v[1] >= c { continue }
                if grid[v[0]][v[1]] == 1 { grid[v[0]][v[1]] = 2; rotten = append(rotten, [2]int{ v[0], v[1] }) }
            }
        }
        if len(rotten) == 0 { break }
        timer++
    }
    for i := 0; i < r; i++ { for j := 0; j < c; j++ { if grid[i][j] == 1 { return -1 } } }
    return timer
}