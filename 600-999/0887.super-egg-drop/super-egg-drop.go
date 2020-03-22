package problem0887

func bionmialCoeff(x, n, k int) int {
	sum := 0
	term := 1
	for i := 1; i <= n && sum < k; i++{
		term *= x - i +1
		term /= i
		sum += term
	}
	return sum
}


func superEggDrop(K int, N int) int {

	// Initialize low and hight as 1st and last floors.
	low := 1
	high := N

	for low < high {
		mid := (low + high )/ 2
		if bionmialCoeff(mid, K, N) < N {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
	
}

func max(a, b int ) int {
	if a > b {
		return a
	}
	return b
}