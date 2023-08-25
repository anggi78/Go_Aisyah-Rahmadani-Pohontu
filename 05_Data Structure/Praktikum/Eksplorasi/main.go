package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	diagonalLeftToRight := matrix[0][0] + matrix[1][1] + matrix[2][2]
	diagonalRightToLeft := matrix[0][2] + matrix[1][1] + matrix[2][0]
	absoluteDifference := int(math.Abs(float64(diagonalLeftToRight - diagonalRightToLeft)))

	fmt.Printf("Diagonal kiri ke kanan: %d\n", diagonalLeftToRight)
	fmt.Printf("Diagonal kanan ke kiri: %d\n", diagonalRightToLeft)
	fmt.Printf("Perbedaan mutlak: %d\n", absoluteDifference)
}
