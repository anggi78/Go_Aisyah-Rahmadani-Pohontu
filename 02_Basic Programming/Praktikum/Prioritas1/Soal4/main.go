package main

import "fmt"

func main() {
		// Mencetak Angka
	for angka := 1; angka < 101; angka++ {
		if (angka%3 == 0) {
			fmt.Print(" Fizz ")
		} else if (angka%5 == 0) {
			fmt.Print(" Buzz ")
		} else if (angka > 101) {
			break
		} else {
			fmt.Print(angka, " ")
		}
	}

}