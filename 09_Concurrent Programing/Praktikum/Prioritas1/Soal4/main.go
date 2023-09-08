package main 

import (
	"fmt"
	"sync"
)

func factorial(n int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	result := 1

	mu.Lock()

	for i := 1; i <= n; i++ {
		result *= i
	}

	defer mu.Unlock()
	fmt.Printf("Faktorial dari %d adalah %d\n", n, result)
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	num := []int{5, 6, 7, 8, 9}

	for _, n := range num {
		wg.Add(1)
		go factorial(n, &wg, &mu)
	}

	wg.Wait()
}