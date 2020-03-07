package problem200

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	ans := 0
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			ans += int(grid[y][x] - '0')
			dfs(grid, x, y, m, n)
		}
	}
	return ans
}

func dfs(grid [][]byte, x, y, m, n int) {
	if x < 0 || y < 0 || x >= n || y >= m || grid[y][x] == '0' {
		return
	}
	grid[y][x] = '0'
	dfs(grid, x+1, y, m, n)
	dfs(grid, x-1, y, m, n)
	dfs(grid, x, y+1, m, n)
	dfs(grid, x, y-1, m, n)
}
