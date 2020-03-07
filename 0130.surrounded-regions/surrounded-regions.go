func solve(board [][]byte)  {

	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])

	for y:= 0; y<m; y++ {
		dfs(board, 0, y)
		dfs(board, n-1, y)
	}

	for x := 0; x < n; x++ {
		dfs(board, x, 0)
		dfs(board, x, m-1)
	}

	var v = map[byte]byte {
		'G':'O',
		'O':'X',
		'X': 'X',
	}
	for y := 0 ; y < m ; y++ {
		for x :=0; x < n ; x++ {
			board[y][x] =v[board[y][x]]
		}
	}

}

// dfs search the graph and mark the '0' to G.
//
func dfs(board [][]byte, x, y int) {

    m := len(board)
	n := len(board[0])
	// border
	if x < 0 || x >= n || y < 0 || y >= m || board[y][x] != 'O' {
		return
	}

	// Mark the 0 on the border t0 'G'
	board[y][x] = 'G'
	// Search continue along the border
	dfs(board, x-1, y)
	dfs(board, x+1, y)
	dfs(board, x, y-1)
	dfs(board, x, y+1)
}