
// Leetcode: Hard
// Tags: DP, TwoPointers
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	left := 0 
	right := n-1
	// Initialize the maximum height left/right. 
	maxLeft = height[left]
	maxRight = height[right]
	// Answer. 
	ans := 0
	// Loop with 2 pointers. 
	for left < right {
		if maxLeft < maxRight {
			ans += maxLeft - height[left]
			maxLeft = max(maxLeft, height[left++])
		} else {
			ans += maxRight - height[right]
			maxRight = max(maxRight, height[r--])
		}
	}
}
 
func min(a,b int) int {
	if a < b {
		return a 
	}
	return b
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}