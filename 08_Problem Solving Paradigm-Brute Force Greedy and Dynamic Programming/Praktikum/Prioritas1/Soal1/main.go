package main

import (
	"fmt"
)

func countBits(n int) []string {
	ans := make([]string, n+1)
	for i := 0; i <= n; i++ {
		binaryStr := fmt.Sprintf("%b", i)
		ans[i] = binaryStr
	}
	return ans
}

func main() {
	n := 2
	result := countBits(n)
	fmt.Println(result)
}