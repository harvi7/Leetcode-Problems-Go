package dfs

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	color := make([]int, n) // 0 = WHITE, 1 = GRAY, 2 = BLACK
	var ans []int

	var dfs func(node int) bool
	dfs = func(node int) bool {
		if color[node] > 0 {
			return color[node] == 2
		}

		color[node] = 1 // mark as GRAY
		for _, nei := range graph[node] {
			if color[node] == 2 {
				continue
			}
			if color[nei] == 1 || !dfs(nei) {
				return false
			}
		}

		color[node] = 2 // mark as BLACK
		return true
	}

	for i := 0; i < n; i++ {
		if dfs(i) {
			ans = append(ans, i)
		}
	}

	return ans
}
