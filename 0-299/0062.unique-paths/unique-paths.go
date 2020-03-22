package problem0062

func uniquePaths(m int, n int) int {

	// out border return 0.
	if m < 0 || n < n {
		return 0
	}

	// create an solution map to store the dp results.
	ans := make([][]int, n+1)
	for i, _ := range ans {
		ans[i] = make([]int, m+1)
	}

	// the start point has 1 way to get itself.
	ans[1][1] = 1

	for y := 1; y <= n; y++ {
		for x := 1; x <= m; x++ {
			if x == 1 && y == 1 {
				continue
			}
			ans[y][x] = ans[y-1][x] + ans[y][x-1]
		}
	}

	return ans[n][m]

}
