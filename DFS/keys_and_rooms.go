package dfs

func canVisitAllRooms(rooms [][]int) bool {
	seen := make([]bool, len(rooms))
	seen[0] = true

	stack := []int{0} // use slice as stack

	for len(stack) > 0 {
		// pop element
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, newKey := range rooms[curr] {
			if !seen[newKey] {
				seen[newKey] = true
				stack = append(stack, newKey)
			}
		}
	}

	for _, visited := range seen {
		if !visited {
			return false
		}
	}
	return true
}
