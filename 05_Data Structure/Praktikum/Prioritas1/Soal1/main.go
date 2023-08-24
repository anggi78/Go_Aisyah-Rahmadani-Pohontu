package main

import "fmt"

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}

func ArrayMerge(arrayA, arrayB []string) []string {
	Merge := arrayA

	for _, item := range arrayB {
		found := false
		for _, existingItem := range arrayA {
			if item == existingItem {
				found = true
				break
			}
		}
		if !found {
			Merge = append(Merge, item)
		}
	}
	return Merge
}