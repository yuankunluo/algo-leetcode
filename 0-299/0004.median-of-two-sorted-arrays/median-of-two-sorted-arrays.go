package problem0004


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	
	len1 := len(nums1)
	len2 := len(nums2)

	// Maker sure n1 <= n2.
	if len1 > len2 {  
		return findMedianSortedArrays(nums2, nums1)
	}
	// nums1 is empty, then the median is in nums2.
	if len1 == 0 {
		return float64(nums2[(len2 -1)/2] + nums2[len2/2]) / 2
	}

	// The total length of the merged array.
	lenM := len1 + len2
	// [startK, endK] is the range for Binary Search.
	startK := 0
	endK := len1

	// The MaxInt/MinInt constants.
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = (0 - MaxInt) -1
	
	// Start the loopping.
	for startK <= endK {
		// Poisitions of the cut position in nums1 and nusm2.
		cut1 := (startK + endK) / 2
		// cut2 is the position of the cut in nums2.
		cut2 := (lenM + 1 )/ 2 - cut1
		// Find l1, l2, r1, r2
		var l1, l2 int 
		
		if cut1 == 0 {
			l1 = MinInt
		} else {
			l1 = nums1[cut1-1]
		}

		if cut2 == 0 {
			l2 = MinInt
		} else {
			l2 = nums2[cut2-1]
		}

		var r1, r2 int

		if cut1 == len1 {
			r1 = MaxInt
		} else {
			r1 = nums1[cut1]
		}

		if cut2 == len2 {
			r2 = MaxInt
		} else {
			r2 = nums2[cut2]
		}

		// The search logic.
		if l1 > r2 {
			endK = cut1 - 1
		} else if l2 > r1 {
			startK = cut1 + 1
		} else {
			if lenM %2 == 0 {
				return float64( (max(l1, l2) + min(r1, r2) ))/2
			}  
			return float64(max(l1, l2))
		}
	}

	return -1

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}