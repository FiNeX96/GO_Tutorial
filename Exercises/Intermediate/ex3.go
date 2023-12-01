package main

import (
    "fmt"
    "strings"
)

// Complete the function below to count the frequency of each word in a string

func WordFrequency(text string) map[string]int {
    // Your code here
}

func main() {
    input := "go is a programming language go is cool go"
    frequencies := WordFrequency(input)
    fmt.Println(frequencies) // Should print map[go:3 is:2 a:1 programming:1 language:1 cool:1]
}