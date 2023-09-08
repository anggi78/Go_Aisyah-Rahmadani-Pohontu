package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = math.MaxInt64
	}

	dp[0] = 0
	for i := 1; i < n; i++ {
		if i-1 >= 0 {
			dp[i] = min(dp[i], dp[i-1]+abs(jumps[i]-jumps[i-1]))
		}
		if i-2 >= 0 {
			dp[i] = min(dp[i], dp[i-2]+abs(jumps[i]-jumps[i-2]))
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}