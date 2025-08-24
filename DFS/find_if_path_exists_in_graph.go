package dfs

var seen bool

func validPath(n int, edges [][]int, source int, destination int) bool {
	visited := make([]bool, n)
	graph := make([]map[int]struct{}, n)

	for i := 0; i < n; i++ {
		graph[i] = make(map[int]struct{})
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u][v] = struct{}{}
		graph[v][u] = struct{}{}
		if _, ok := graph[source][destination]; ok { // direct connection exists
			return true
		}
	}

	seen = false
	dfsValidPath(graph, visited, source, destination)
	return seen
}

func dfsValidPath(graph []map[int]struct{}, visited []bool, start, end int) {
	if !visited[start] && !seen {
		if start == end {
			seen = true
			return
		}

		visited[start] = true
		for neighbor := range graph[start] {
			dfsValidPath(graph, visited, neighbor, end)
		}
	}
}
