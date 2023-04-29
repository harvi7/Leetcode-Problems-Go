func integerBreak(n int) int {
    dp := make([]int, n + 1)
    for i := 0; i < n + 1; i++ {
        for j := 0; j < i + 1; j++ {
            dp[i] = max(dp[i], max(dp[j] * (i - j), (i - j) * j))
        }
    }
    return dp[n]
}

func max(a, b int) int {
    if a > b { return a }
    return b
}