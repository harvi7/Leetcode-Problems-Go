func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    firstBuy, firstSell, secondBuy, secondSell := math.MinInt64, 0, math.MinInt64, 0
    for i := 0; i < len(prices); i++ {
        firstBuy = max(firstBuy, -prices[i])
        firstSell = max(firstSell, firstBuy + prices[i])
        secondBuy = max(secondBuy, firstSell - prices[i])
        secondSell = max(secondSell, secondBuy + prices[i])
    }
    return max(0, secondSell)
}

func max(x int, y int) int {
    if x > y {
        return x
    }
    return y
}