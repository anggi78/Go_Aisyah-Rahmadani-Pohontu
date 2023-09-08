package main

import "fmt"

func printMultiples3(c chan int) {
	for i := 1; i <= 10; i++ {
		if i%3 == 0 {
			c <- i
		}
	}
	close(c)
}

func main() {
	bufferSize := 3

	ch := make(chan int, bufferSize)

	go printMultiples3(ch)

	for num := range ch {
		fmt.Println(num)
	}
}