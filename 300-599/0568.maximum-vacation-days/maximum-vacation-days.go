package problem0568

func maxVacationDays(flights [][]int, days [][]int) int {
	// n is the number of cities.
	n := len(flights)
	// k is the number of weeks.
	k := len(days[0])
	// compute the MinInt value
	const MaxInt = int(^uint(0) >> 1)
	const MinInt = -MaxInt - 1
	// dp is a one-d slice to record the current longest days
	// for the city dp[i]. It initializes with -1, but
	// the start city is 0.
	dp := make([]int, n)
	for i := range dp {
		dp[i] = MinInt
	}
	dp[0] = 0

	// Loop though weeks, there r K weeks.
	for w := 0; w < k; w++ {
		// Make a temporary slice to record the
		// computing results.
		temp := make([]int, n)
		for i := range temp {
			temp[i] = MinInt
		}
		// Loop though the flight matrix.
		// d stands for destination
		// s stands for source
		for d := 0; d < n; d++ {
			for s := 0; s < n; s++ {
				if d == s || flights[s][d] == 1 {
					temp[d] = max(temp[d], dp[s]+days[d][w])
				}
			}
		}

		// Update the final results
		dp = append(temp[:0:0], temp...)
	}
	return findMax(dp)
}

// max is a helper function to get the max value from
// two integer.
func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// findMax is helper function to find the maximum value in
// a slice.
func findMax(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
