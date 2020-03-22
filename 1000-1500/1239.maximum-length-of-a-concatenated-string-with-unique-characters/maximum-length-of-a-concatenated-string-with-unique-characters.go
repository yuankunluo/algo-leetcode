package problem1239

func maxLength(arr []string) int {
	mask := 0
	return dfs(arr, 0, mask)
}

func isUnique(s string, mask *int) bool {
	for _, c := range s {
		i := c - 'a'
		if *mask&(1<<i) != 0 {
			return false
		}
		*mask |= 1 << i
	}
	return true
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func dfs(arr []string, childIndex int, mask int) int {
	// Already reach the end.
	if childIndex == len(arr) {
		return 0
	}

	// A temproray bitmask
	mask2 := mask
	// If this string is unique
	if isUnique(arr[childIndex], &mask2) {
		curLen := len(arr[childIndex])
		len1 := curLen + dfs(arr, childIndex+1, mask2)
		len2 := dfs(arr, childIndex+1, mask)
		return max(len1, len2)
	}

	return dfs(arr, childIndex+1, mask)

}
