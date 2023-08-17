package main

import "fmt"

func main() {
	// Luas Trapesium
	var alas int
	var atas int
	var tinggi int
	var Luas int

    fmt.Print("Masukkan nilai alas: ")
    fmt.Scan(&alas)

	fmt.Print("Masukkan nilai atas: ")
    fmt.Scan(&atas)

	fmt.Print("Masukkan nilai tinggi: ")
    fmt.Scan(&tinggi)

	Luas = (alas + atas) / 2 * tinggi
	fmt.Println("Luas:", Luas)
}