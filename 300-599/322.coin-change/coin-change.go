package problem0322

// LeetCode: Medium
// Tags: DP
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] >= amount+1 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
