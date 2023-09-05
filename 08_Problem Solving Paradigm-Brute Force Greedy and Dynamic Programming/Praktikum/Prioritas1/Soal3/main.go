package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	fib0, fib1 := 0, 1
	result := 0

	for i := 2; i <= n; i++ {
		result = fib0 + fib1
		fib0, fib1 = fib1, result
	}

	return result
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}