package graph

func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	dfsAllPaths(graph, 0, []int{}, &res)
	return res
}

func dfsAllPaths(graph [][]int, node int, path []int, res *[][]int) {
	path = append(path, node)
	if node == len(graph)-1 {
		// Make a copy of path and append to result
		*res = append(*res, append([]int{}, path...))
		return
	}

	for _, v := range graph[node] {
		dfsAllPaths(graph, v, path, res)
	}
}
