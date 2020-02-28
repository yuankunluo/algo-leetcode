package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// f[i] is the longset continuous incresing subsequence's
	// length till i.
	f := make([]int, len(nums))
	// The start.
	f[0] = 1
	ans := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			// Accumulator.
			f[i] = f[i-1] + 1
			if f[i] > ans {
				ans = f[i]
			}
		} else {
			// Reset.
			f[i] = 1
		}
	}
	return ans
}
