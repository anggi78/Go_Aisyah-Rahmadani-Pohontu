package main

import "fmt"

func fibonacci(number int) int {
	fib := make([]int, number+1)

	fib[0] = 0
	if number > 0 {
		fib[1] = 1
	}

	for i := 2; i <= number; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	
	return fib[number]
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}