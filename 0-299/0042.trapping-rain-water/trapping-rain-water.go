
// Leetcode: Hard
// Tags: DP, TwoPointers
func trap(height []int) int {
	// Get the input height.
	n := len(height)
	ans := 0
	// left[i] stores the maximum height from 0 to i. 
	left := make([]int, n)
	// right[i] stores the maximum height from i to n-1;
	right := make([]int, n)
	for i :=0; i < n; i++ {
		if i == 0 {
			left[i] = height[i]
		} else {
			left[i] = max(left[i-1], height[i])
		}
	}
	for i := n-1; i >= 0; i-- {
		if i == n -1  {
			right[i] = height[i]
		} else {
			right[i] = max(right[i+1], height[i])
		}
	}
	// Compute the ans. 
	for i := 0; i< n; i++ {
		ans += min(left[i], right[i]) - height[i]
	}
	return ans;
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