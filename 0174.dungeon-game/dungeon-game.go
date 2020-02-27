package problem0174

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])




	// Initialize a dp 2d-array with MaxInt32
	// 
	const MaxUint32 = ^uint32(0)
	const MaxInt32 = int(MaxUint32 >> 1)
	
	dp := make([]int, n + 1)
	for i := range dp {
		dp[i] = MaxInt32
	}
	dp[n-1] = 1

	for y := m-1; y >= 0; y-- {
		for x := n-1; x >=0; x-- {
			dp[x] = max(1, min(dp[x], dp[x + 1]) - dungeon[y][x])
		}
	}
	

	

	return dp[0]
}

func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}


