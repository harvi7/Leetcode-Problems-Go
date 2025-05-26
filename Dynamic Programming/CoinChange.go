func coinChange(coins []int, amount int) int {
    // dp  := make([]int, amount + 1)
    // for i := 1; i <= amount; i++ {
    //     dp[i] = -1
    //     for _, c := range coins {
    //         if i >= c && dp[i - c] != -1 && (dp[i] == -1 || dp[i - c] + 1 < dp[i]) {
    //             dp[i] = dp[i - c] + 1
    //         }     
    //     }
    // }
    
    // return dp[amount]

    max := amount + 1;
    dp := make([]int, amount + 1)
    dp[0] = 0
    for i := 1; i <= amount; i++ {
        dp[i] = max
        for _, c := range coins {
            if (c <= i) {
                dp[i] = min(dp[i], dp[i - c] + 1);
            }
        }
    }
    if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}