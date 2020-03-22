package main

import (
	"fmt"
)

type NumArray struct {
	Sums []int
}

func Constructor(nums []int) NumArray {

	numarray := NumArray{
		Sums: make([]int, len(nums)),
	}

	if len(nums) == 0 {
		return numarray
	}

	numarray.Sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		numarray.Sums[i] += nums[i] + numarray.Sums[i-1]
	}
	return numarray
}

func (this *NumArray) SumRange(i int, j int) int {

	if i == 0 {
		return this.Sums[j]
	}

	return this.Sums[j] - this.Sums[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	param_1 := obj.SumRange(0, 2)
	fmt.Println(param_1)
}
