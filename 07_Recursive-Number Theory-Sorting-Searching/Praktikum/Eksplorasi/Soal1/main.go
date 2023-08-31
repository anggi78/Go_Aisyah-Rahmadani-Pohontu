package main

import (
	"fmt"
)

func MaxSequence(arr []int) int {
	maxSum := 0
	n := len(arr)

	for i := 0; i < n; i++ {
		currSum := 0
		for j := i; j < n; j++ {
			currSum += arr[j]
			if currSum > maxSum {
				maxSum = currSum
			}
		}
	}
	return maxSum
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
}
