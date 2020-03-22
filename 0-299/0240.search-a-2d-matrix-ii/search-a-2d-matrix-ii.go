package main

func searchMatrix(matrix [][]int, target int) bool {

	y := len(matrix)

	if y == 0 {
		return false
	}
	x := len(matrix[0])

	i := 0
	j := x - 1

	for i < y && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}
