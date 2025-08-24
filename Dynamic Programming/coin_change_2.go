package dp

// https://leetcode.com/problems/coin-change-ii/solutions/3226141/golang-dp/
// https://www.youtube.com/watch?v=ruMqWViJ2_U&ab_channel=Techdose

func change(amount int, coins []int) int {

	// n := len(coins)
	// // Create DP table: dp[i][j] = ways to form amount j using first i coins
	// dp := make([][]int, n+1)
	// for i := range dp {
	// 	dp[i] = make([]int, amount+1)
	// }

	// // Fill first column: amount 0 can always be formed in exactly 1 way (no coins)
	// for i := 0; i <= n; i++ {
	// 	dp[i][0] = 1
	// }

	// // Fill DP table
	// for i := 1; i <= n; i++ {
	// 	for j := 1; j <= amount; j++ {
	// 		if j >= coins[i-1] {
	// 			// Include coin[i-1] + Exclude coin[i-1]
	// 			dp[i][j] = dp[i][j-coins[i-1]] + dp[i-1][j]
	// 		} else {
	// 			// Can't include this coin
	// 			dp[i][j] = dp[i-1][j]
	// 		}
	// 	}
	// }

	// return dp[n][amount]

	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
