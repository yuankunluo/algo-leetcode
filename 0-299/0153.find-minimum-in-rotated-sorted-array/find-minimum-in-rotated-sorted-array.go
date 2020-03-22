package problem0153

func findMin(nums []int) int {
	return findMinRec(nums, 0, len(nums)-1)
}

// findMinRec will find the minimum in a subarray range from l to r.
//
//
func findMinRec(nums []int, l, r int) int {
	// Only has 1 element.
	if l == r {
		return nums[l]
	}

	// Only has 2 elements.
	if l+1 == r {
		return min(nums[l], nums[r])
	}

	// The subarray is sorted, return the left element.
	if nums[l] < nums[r] {
		return nums[l]
	}

	// divide and conquer
	mid := l + (r-l)/2
	return min(findMinRec(nums, l, mid-1), findMinRec(nums, mid, r))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
