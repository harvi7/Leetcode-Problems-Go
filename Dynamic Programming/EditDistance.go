package dp

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 0
			} else if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + minNums(dp[i-1][j-1], dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[m][n]
}

func minNums(nums ...int) int {
	ans := nums[0]
	for _, v := range nums {
		if v < ans {
			ans = v
		}
	}
	return ans
}
