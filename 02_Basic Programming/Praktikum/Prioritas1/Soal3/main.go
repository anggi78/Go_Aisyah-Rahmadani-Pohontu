package main

import "fmt"

func main() {
	// Grade
	var grade int
    fmt.Print("Masukkan nilai: ")
    fmt.Scan(&grade)
	switch {
	case grade >= 80 && grade <= 100 :
		fmt.Printf("Kamu dapat nilai A")
	case grade >= 65 && grade <= 79 :
		fmt.Printf("Kamu dapat nilai B")
	case grade >= 50 && grade <= 64 :
		fmt.Printf("Kamu dapat nilai C")
	case grade >= 35 && grade <= 49 :
		fmt.Printf("Kamu dapat nilai D")
	case grade >= 0 && grade <= 34 :
		fmt.Printf("Kamu dapat nilai E")
	default:
		fmt.Printf("Invalid Nilai")
	}

}