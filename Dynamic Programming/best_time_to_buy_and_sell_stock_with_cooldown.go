func maxProfit(prices []int) int {
    sold := math.MinInt32
    held := math.MinInt32
    reset := 0
    
    for _, price := range prices {
        pre_sold := sold
        sold = held + price
        held = max(held, reset - price)
        reset = max(reset, pre_sold)
    }
    return max(reset, sold)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
