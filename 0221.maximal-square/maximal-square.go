package main

// TODO: continue
func maximalSquare(matrix [][]byte) int {

	// The empty case.
	if len(matrix) == 0 {
		return 0
	}
	// Got the dimensions.
	height := len(matrix)
	width := len(matrix[0])
	// Construct a DP-2D.
	sizes := make([][]int, height)
	for i := range sizes {
		sizes[i] = make([]int, width)
	}
	// Answer.
	ans := 0
	// Loop though matrix and compute.
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			sizes[i][j] = matrix[i][j] - '0'
			
			if sizes[i][j] 
		}
	}



}

func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}