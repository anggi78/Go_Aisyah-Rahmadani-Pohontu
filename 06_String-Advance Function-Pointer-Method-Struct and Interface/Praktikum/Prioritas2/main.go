package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(caesar(3, "abc")) // def
	fmt.Println(caesar(2, "alta")) // cnvc
	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}

func caesar(offset int, input string) string {
	result := ""
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, char := range input {
		if char == ' ' {
			result += " " 
			continue
		}

		charIndex := strings.IndexRune(alphabet, char)
		if charIndex == -1 {
			result += string(char)
		} else {
			newIndex := (charIndex + offset) % len(alphabet)
			result += string(alphabet[newIndex])
		}
	}
	return result
}
