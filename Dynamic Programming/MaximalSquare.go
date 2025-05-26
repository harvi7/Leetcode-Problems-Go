func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    dp := make([][]int, m)

    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }

    max := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j]=='1'{
                max = 1
                dp[i][j] = 1
            } 
        }
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if dp[i][j] == 1 {
                dp[i][j] = Min(dp[i - 1][j - 1], Min(dp[i - 1][j], dp[i][j - 1])) + 1
                if max < dp[i][j ]{
                    max = dp[i][j]
                }
            }
        }
    }
    return max*max
}

func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}