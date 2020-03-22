package main

func majorityElement(nums []int) int {
	majority := nums[0]
	count := 0

	for _, v := range nums {
		if v == majority {
			count++
		} else {
			count--
			if count == 0 {
				count = 1
				majority = v
			}

		}

	}
	return majority
}
