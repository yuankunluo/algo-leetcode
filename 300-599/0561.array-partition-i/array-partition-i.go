package main

func arrayPairSum(nums []int) int {

	// The maximum of inputs' range.
	const maxValue = 10000
	// The count table. Since the range is [-10000,10000],
	// we will shift every value right 10000.
	count := make([]int, 2*maxValue+1)

	// Do counting.
	for _, v := range nums {
		count[v+maxValue]++
	}

	ans := 0
	n := -maxValue
	first := true

	for n <= maxValue {
		if count[n+maxValue] == 0 {
			n++
			continue
		}
		if first {
			ans += n
			first = false
		} else {
			first = true
		}
		count[n+maxValue]--
	}

	return ans
}
