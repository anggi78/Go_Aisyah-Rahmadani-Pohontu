package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Masukkan kata atau frase: ")
	fmt.Scan(&input)
	fmt.Println("Captured: ", input)

	if isPalindrome(input) {
		fmt.Println("Kata tersebut adalah palindrom.")
	} else {
		fmt.Println("Kata tersebut bukan palindrom.")
	}
}

func isPalindrome(str string) bool {
	str = strings.ToLower(strings.ReplaceAll(str, " ", ""))
	left := 0
	right := len(str) - 1

	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}

	return true
}
