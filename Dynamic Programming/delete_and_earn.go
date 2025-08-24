package dp

func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	tmp := make([]int, 10001)
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
		tmp[v] += v
	}
	dp := make([]int, max+1)
	dp[0] = tmp[0]
	dp[1] = tmp[1]
	for i := 2; i < len(dp); i++ {
		dp[i] = Max(dp[i-1], dp[i-2]+tmp[i])
	}
	return dp[max]
}
