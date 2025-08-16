package dfs

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	graph := make(map[string][]string) // adjacency list
	emailToName := make(map[string]string)

	buildGraph(graph, emailToName, accounts)

	var res [][]string
	visited := make(map[string]bool)

	for email := range emailToName {
		if !visited[email] {
			var list []string
			dfsAccountsMerge(graph, email, &list, visited)
			sort.Strings(list)
			// Add the name at the front
			name := emailToName[email]
			list = append([]string{name}, list...)
			res = append(res, list)
		}
	}
	return res
}

func buildGraph(graph map[string][]string, emailToName map[string]string, accounts [][]string) {
	for _, acc := range accounts {
		name := acc[0]
		for i := 1; i < len(acc); i++ {
			email := acc[i]
			emailToName[email] = name

			if i == 1 {
				continue
			}
			prev := acc[i-1]
			// Add both directions
			graph[prev] = append(graph[prev], email)
			graph[email] = append(graph[email], prev)
		}
	}
}

func dfsAccountsMerge(graph map[string][]string, email string, list *[]string, visited map[string]bool) {
	visited[email] = true
	*list = append(*list, email)

	for _, neighbor := range graph[email] {
		if !visited[neighbor] {
			dfsAccountsMerge(graph, neighbor, list, visited)
		}
	}
}
