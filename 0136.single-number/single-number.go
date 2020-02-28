package main

func singleNumber(nums []int) int {
	x := 0
	// XOR computation
	for i := range nums {
		x ^= nums[i]
	}
	return x
}
