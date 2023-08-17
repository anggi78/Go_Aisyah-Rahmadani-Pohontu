package main

import "fmt"

func main() {
		// Mencetak Angka
	for angka := 1; angka < 101; angka++ {
		if (angka%3 == 0) {
			fmt.Println("Fizz")
		} else if (angka%5 == 0) {
			fmt.Println("Buzz")
		} else if (angka > 101) {
			break
		} else {
			fmt.Println(angka)
		}
	}

}