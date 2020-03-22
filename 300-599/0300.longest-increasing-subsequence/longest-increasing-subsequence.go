package problem0300

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// create a dp to remember the lts at position i
	dp := make([]int, len(nums))
	for i,_ := range(dp){
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j]{
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}
	return findMax(dp)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func findMax(nums []int) int {
	max := nums[0]
	for _, v := range(nums){
		if v > max {
			max =  v
		}
	}
	return max
}