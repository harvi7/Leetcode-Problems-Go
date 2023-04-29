func maxProfit(prices []int, fee int) int {
    cash := 0 
    hold := -prices[0]
    for i := 1; i < len(prices); i++ {
        cash = Max(cash, hold + prices[i] - fee)
        hold = Max(hold, cash - prices[i])
    }
    return cash
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}