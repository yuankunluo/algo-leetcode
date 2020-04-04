// Leetcode: Medium
// Tags: Prefix Sum, Hashtable
// Runtime: 24ms
// Ram: 6.3 MB
func subarraySum(nums []int, k int) int {

	if nums == nil || len(nums) == 0 {
		return 0
	}

	sums := make(map[int]int);
	sums[0] = 1
	prefixSum := 0
	count := 0 
	for _, v := range nums {
		prefixSum += v 
		count += sums[prefixSum - k]
		sums[prefixSum] = sums[prefixSum] + 1
	}
	return count

}
// A better solution.
func subarraySum_2(nums []int, k int) int {
	res, sum := 0, 0
	rec := make(map[int]int, len(nums))
	rec[0]  = 1
	for i := range nums {
		sum += nums[i]
		res += rec[sum-k]
		rec[sum]++
	}
	return res
}

