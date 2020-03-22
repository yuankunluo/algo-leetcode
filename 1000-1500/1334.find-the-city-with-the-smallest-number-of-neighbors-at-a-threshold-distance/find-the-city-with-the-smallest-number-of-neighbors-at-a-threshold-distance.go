package pboblem1334

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// The maxInt construct.
	const MaxUint = ^uint32(0)
	const MaxInt = int(int32(MaxUint>>1)) / 2

	// Construct DP 2-D array for Floyd Warshall.
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = MaxInt
		}
	}

	// Fill up the DP with edges
	for _, e := range edges {
		dp[e[0]][e[1]] = e[2]
		dp[e[1]][e[0]] = e[2]
	}

	// Floyd-Washall Algorithm to compute the single min path.
	for k := 0; k < n; k++ {
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				dp[u][v] = min(dp[u][v], dp[u][k]+dp[k][v])
			}
		}
	}

	ans := -1
	minNb := MaxInt
	for u := 0; u < n; u++ {
		nb := 0
		for v := 0; v < n; v++ {
			if v != u && dp[u][v] <= distanceThreshold {
				nb++
			}
		}
		if nb <= minNb {
			minNb = nb
			ans = u
		}
	}

	return ans

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
