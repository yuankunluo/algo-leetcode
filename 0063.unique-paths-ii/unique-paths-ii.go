package problem0063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// the start point is obstacle
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	// out border return 0.
	if m < 0 || n < 0 {
		return 0
	}

	// create an solution map to store the dp results.
	ans := make([][]int, n+1)
	for i := range ans {
		ans[i] = make([]int, m+1)
	}

	// the start point has 1 way to get itself.
	ans[1][1] = 1

	for y := 1; y <= n; y++ {
		for x := 1; x <= m; x++ {
			// start point
			if x == 1 && y == 1 {
				continue
			}

			if obstacleGrid[y-1][x-1] == 1 {
				ans[y][x] = 0
			} else {
				ans[y][x] = ans[y-1][x] + ans[y][x-1]
			}
		}
	}

	return ans[n][m]

}
