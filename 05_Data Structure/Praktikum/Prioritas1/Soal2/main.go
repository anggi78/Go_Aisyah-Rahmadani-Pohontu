package main

import "fmt"

func main() {
	// Test cases
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) // map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))                      // map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))                                         // map[]
}

func Mapping(slice []string) map[string]int {
	countMap := make(map[string]int)

	for _, s := range slice {
		countMap[s]++
	}

	return countMap
}