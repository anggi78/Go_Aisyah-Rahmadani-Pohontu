package main

import (
	"fmt"
)

func main() {
	fmt.Println(prima(1000000007)) //true
	fmt.Println(prima(13)) //true
	fmt.Println(prima(17)) //true
	fmt.Println(prima(20)) //false
	fmt.Println(prima(35)) //false
	fmt.Println(prima(2)) //true
	fmt.Println(prima(2000000007)) //false
	fmt.Println(prima(97)) //true
}

func prima(number int) bool {
	for i := 2; i*i <= number; i++ {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}
	return true
}