package problem0277

// LeetCode: Medium
/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		// 1. Find a candidate by one pass.
		candidate := 0
		for i := 1; i < n; i++ {
			if !knows(i, candidate) {
				candidate = i
			}
		}
		// 2. Make sure if the candidate is a celebrity.
		for i := 0; i < n; i++ {
			if i == candidate {
				continue
			}
			if !knows(i, candidate) || knows(candidate, i) {
				return -1
			}

		}
		return candidate
	}
}
