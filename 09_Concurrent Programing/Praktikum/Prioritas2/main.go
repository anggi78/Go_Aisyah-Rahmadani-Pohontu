package main

import (
    "fmt"
    "strings"
    "sync"
)

func main() {
    text := "Lorem ipsum dolor sit amet consectetur adipiscing elit"

    var wg sync.WaitGroup

    wg.Add(1)
    go frekuensi(text, &wg)
    wg.Wait()
}

func frekuensi(text string, wg *sync.WaitGroup) {
    defer wg.Done()
    newText := strings.Replace(text, " ", "", -1)

    frequencyMap := make(map[rune]int)
    for _, char := range newText {
        count := frequencyMap[char]
        frequencyMap[char] = count + 1
    }

    for char, count := range frequencyMap {
        fmt.Printf("%c: %d\n", char, count)
    }
}