package main

func matrixReshape(nums [][]int, r int, c int) [][]int {

	if len(nums) == 0 {
		return nums
	}

	m := len(nums)
	n := len(nums[0])

	if m*n != r*c {
		return nums
	}

	ans := make([][]int, r)
	for i := range ans {
		ans[i] = make([]int, c)
	}

	for i := 0; i < m*n; i++ {
		oldY := i / n
		oldX := i % n

		newY := i / c
		newX := i % c

		ans[newY][newX] = nums[oldY][oldX]
	}
	return ans
}
