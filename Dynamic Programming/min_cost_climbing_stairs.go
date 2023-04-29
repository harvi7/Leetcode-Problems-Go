func minCostClimbingStairs(cost []int) int {
    step1, step2 := 0, 0
    for i := len(cost) - 1; i >= 0; i-- {
		step1, step2 = step2, cost[i] + min(step1, step2)
	}
    return min(step1, step2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}