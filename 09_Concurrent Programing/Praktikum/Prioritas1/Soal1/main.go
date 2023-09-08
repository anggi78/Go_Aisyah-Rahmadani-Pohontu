package main 

import (
	"fmt"
	"time"
)

func printMultiples(x int) {
	for i := 1; ; i++ {
		if i%x == 0 {
			fmt.Printf("%d is a multiple of %d\n", i, x)
		}
		time.Sleep(3*time.Second)
	}
}

func main() {
	x := 3

	go printMultiples(x)

	select{}
}