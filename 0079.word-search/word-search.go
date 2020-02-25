func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}

	// get the width and height of board.
	h := len(board)
	w := len(board[0])

	for i:= 0 ; i < w; i++ {
		for j:= 0 ; j < h; j++ {
			if search(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

// search a word in the board from position (x,y).
//
//
func search(board [][]byte, word string, d, x, y int) bool {
	// out of boarder or can not find the character.
	if x < 0 || x == len(board[0]) || y <0 || y == len(board) || word[d] != board[y][x]{
		return false
	}

	// find the last char of the word.
	if d == len(word) -1 {
		return true
	}

	// recursively call.
	curChar := board[y][x]
	board[y][x] = 0
	found := search(board, word, d+1, x+1, y) || search(board, word, d+1, x-1, y)|| search(board, word, d+1, x, y+1)|| search(board, word, d+1, x, y-1)
	board[y][x] = curChar
	return found
}