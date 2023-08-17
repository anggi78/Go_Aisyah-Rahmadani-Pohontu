package main

import "fmt"

func main() {
	// Segitiga Asterik
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}