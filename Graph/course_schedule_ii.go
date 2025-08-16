package graph

func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegree := make([]int, numCourses)
	graph := make([][]int, numCourses)

	// Build graph & count in-degrees
	for _, pre := range prerequisites {
		course, prereq := pre[0], pre[1]
		graph[prereq] = append(graph[prereq], course)
		inDegree[course]++
	}

	// Stack (or queue) for courses with no prerequisites
	stack := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			stack = append(stack, i)
		}
	}

	order := []int{}

	// Process stack
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		order = append(order, curr)

		// Decrease in-degree of dependent courses
		for _, next := range graph[curr] {
			inDegree[next]--
			if inDegree[next] == 0 {
				stack = append(stack, next)
			}
		}
	}

	// If we couldn't take all courses, return empty
	if len(order) != numCourses {
		return []int{}
	}

	return order
}
