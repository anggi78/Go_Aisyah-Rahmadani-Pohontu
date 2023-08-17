package main

import "fmt"

func main() {
	// Ganjil Genap
	var number int

	fmt.Printf("Masukkan angka: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("Bilangan Genap")
	} else {
		fmt.Println("Bilangan Ganjil")
	}

}