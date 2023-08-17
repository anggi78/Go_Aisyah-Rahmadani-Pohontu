package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan kata: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Println("Captured:", input)

	if isPalindrome(input) {
		fmt.Println("Palindrom.")
	} else {
		fmt.Println("Bukan Palindrom.")
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
