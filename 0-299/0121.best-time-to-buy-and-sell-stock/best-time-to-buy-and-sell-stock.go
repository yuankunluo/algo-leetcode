// DP

func maxProfit(prices []int) int {
	n := len(prices)
	// Have less than 1 days.
	if n < 1 {
		return 0
	}
	// Create two slices to record the minPrices to buy till ith,
	// and max profit till ith day to sell.
	minPrices := make([]int, n)
	maxProfit := make([]int, n)
	// The 0th day to buy is the same as the 0th day stock price.
	minPrices[0] = prices[0]
	// Can not sell stock on 0th day.
	maxProfit[0] = 0
	for i := 1; i < n ; i++ {
		minPrices[i] = min(minPrices[i-1], prices[i])
		maxProfit[i] = max(maxProfit[i-1], prices[i] - minPrices[i-1])
	}
	return maxProfit[n-1]
	
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}