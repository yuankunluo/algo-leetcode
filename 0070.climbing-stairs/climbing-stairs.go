package main

import "fmt"

func climbStairs(n int) int {
	// f[i] = climbStairs(i)
	f := make([]int, n+1)

	// Some important initializing.
	// 1 way to climb 0 step stair.
	f[0] = 1
	// 1 way to climb 1 step stair.
	f[1] = 1

	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	return f[n]
}

func main() {
	ans := climbStairs(3)
	fmt.Println(ans)
}
