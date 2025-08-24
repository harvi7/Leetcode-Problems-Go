package dp

func maxProfitIV(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}

	if k >= len(prices) {
		return allTimeProfit(prices)
	}

	T := make([]int, len(prices))
	prev := make([]int, len(prices))

	for i := 1; i <= k; i++ {
		maxDiff := -prices[0]
		for j := 1; j < len(prices); j++ {
			if maxDiff+prices[j] > T[j-1] {
				T[j] = maxDiff + prices[j]
			} else {
				T[j] = T[j-1]
			}
			if prev[j]-prices[j] > maxDiff {
				maxDiff = prev[j] - prices[j]
			}
		}
		for j := 1; j < len(prices); j++ {
			prev[j] = T[j]
		}
	}

	return T[len(T)-1]
}

func allTimeProfit(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	profit := 0
	localMin := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i-1] >= arr[i] {
			localMin = arr[i]
		} else {
			profit += arr[i] - localMin
			localMin = arr[i]
		}
	}

	return profit
}
