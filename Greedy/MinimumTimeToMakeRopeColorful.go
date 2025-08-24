package greedy

func minCost(colors string, neededTime []int) int {
	totalTime, currMaxTime := 0, 0

	for i := 0; i < len(colors); i++ {
		if i > 0 && colors[i] != colors[i-1] {
			currMaxTime = 0
		}

		totalTime += min(currMaxTime, neededTime[i])
		currMaxTime = max(currMaxTime, neededTime[i])
	}

	return totalTime
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
