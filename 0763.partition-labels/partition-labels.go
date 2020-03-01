package main

func partitionLabels(S string) []int {
	// Use a 128 length slice to
	// store all leters' frequency, ascii code has 128 slots.
	lastIndex := make([]int, 128)
	for i, c := range S {
		lastIndex[int(c)] = i
	}
	ans := make([]int, 0)
	start := 0
	end := 0
	for i := 0; i < len(S); i++ {
		end = max(lastIndex[int(S[i])], end)
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
