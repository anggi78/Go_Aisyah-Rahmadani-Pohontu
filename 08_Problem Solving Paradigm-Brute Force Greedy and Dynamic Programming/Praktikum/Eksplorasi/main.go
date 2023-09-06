package main

import "fmt"

func numToRomawi(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	rom := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romawiNum := ""
	i := 0

	for num > 0 {
		k := num / val[i]
		for j := 0; j < k; j++ {
			romawiNum += rom[i]
			num -= val[i]
		}
		i++
	}
	return romawiNum
}

func main() {
	testCases := []int{4, 9, 23, 2021, 1646}
    for _, num := range testCases {
        roman := numToRomawi(num)
        fmt.Printf("Input: %d, Output: %s\n", num, roman)
    }
}