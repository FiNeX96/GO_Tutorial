package main

import (
    "fmt"
    "strings"
)

func WordFrequency(text string) map[string]int {
    words := strings.Fields(text)
    frequency := make(map[string]int)

    for _, word := range words {
        frequency[word]++
    }

    return frequency
}

func main() {
    input := "go is a programming language go is cool go"
    frequencies := WordFrequency(input)
    fmt.Println(frequencies) // Prints map[a:1 cool:1 go:3 is:2 language:1 programming:1]
}
