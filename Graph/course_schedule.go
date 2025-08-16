package graph

// https://youtu.be/rG2-_lgcZzo?si=JIhfSKfiGzDdcNwo

func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make([]int, numCourses)
	count := 0

	// Count in-degrees
	for _, pre := range prerequisites {
		inDegree[pre[0]]++
	}

	// Use slice as a stack
	stack := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			stack = append(stack, i)
		}
	}

	// Process stack
	for len(stack) > 0 {
		// Pop
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++

		// Reduce in-degree for dependent courses
		for _, pre := range prerequisites {
			if pre[1] == curr {
				inDegree[pre[0]]--
				if inDegree[pre[0]] == 0 {
					stack = append(stack, pre[0])
				}
			}
		}
	}

	return count == numCourses
}
