func maxProfit(prices []int) int {
    minprice, maxprofit := math.MaxInt64, 0
    for i := 0; i < len(prices); i++ {
        if prices[i] < minprice {
            minprice = prices[i]
        }  else if prices[i] - minprice > maxprofit {
            maxprofit = prices[i] - minprice
        }          
    }  
    return maxprofit
}