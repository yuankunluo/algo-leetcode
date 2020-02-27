package problem0005

func longestPalindrome(s string) string {
	maxLen := 0
	start := 0
	for i := 0; i < len(s); i++ {
		curMax := max(getMaxLen(i, i, &s), getMaxLen(i, i+1, &s))
		if curMax > maxLen {
			maxLen = curMax
			start = i - (maxLen-1)/2
		}
	}

	return string(s[start : start+maxLen])
}

func getMaxLen(left, right int, s *string) int {
	for left >= 0 && right < len(*s) && (*s)[left] == (*s)[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
