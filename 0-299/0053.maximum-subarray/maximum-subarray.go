package main

import "fmt"

func maxSubArray(nums []int) int {
	// f[i] = f[i-1] > 0 ? nums[i] + f[i-1] : nums[i];
	f := make([]int, len(nums))
	f[0] = nums[0]
	ans := f[0]
	for i := 1; i < len(nums); i++ {
		if f[i-1] > 0 {
			f[i] = f[i-1] + nums[i]
		} else {
			f[i] = nums[i]
		}
		if f[i] > ans {
			ans = f[i]
		}
	}
	return ans

}

func main() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	output := 6
	fmt.Println(output == maxSubArray(input))
}
