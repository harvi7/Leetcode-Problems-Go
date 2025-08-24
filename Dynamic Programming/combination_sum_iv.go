package dp

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for comb_sum := 0; comb_sum < target+1; comb_sum++ {
		for _, num := range nums {
			if comb_sum-num >= 0 {
				dp[comb_sum] += dp[comb_sum-num]
			}
		}
	}
	return dp[target]
}
