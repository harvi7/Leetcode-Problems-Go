func maxScoreSightseeingPair(values []int) int {
    maxi, ans := values[0] + 0, 1
    n := len(values)
    for i := 1; i < n; i++ {
        ans = max(ans, maxi + values[i] - i)
        maxi = max(maxi, values[i] + i)
    }
    return ans
}

func max(a, b int) int {
    if a > b { return a }
    return b
}