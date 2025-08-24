package dp

func mincostTickets(days []int, costs []int) int {
	n := len(days)
	dp := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		d7, d30 := i, i
		for d7 < n && days[d7] < days[i]+7 {
			d7 = d7 + 1
		}
		for d30 < n && days[d30] < days[i]+30 {
			d30 = d30 + 1
		}
		dp[i] = min(costs[0]+dp[i+1], min(costs[1]+dp[d7], costs[2]+dp[d30]))
	}
	return dp[0]
}
