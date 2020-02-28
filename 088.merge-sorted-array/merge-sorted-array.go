package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	first := m - 1
	second := n - 1

	for i := m + n - 1; i >= 0; i-- {
		// nums2 has already been merged.
		if second < 0 {
			break
		}
		if nums1[first] > nums2[second] {
			nums1[i] = nums1[first]
			first--
		} else {
			nums1[i] = nums2[second]
			second--
		}
	}
}
