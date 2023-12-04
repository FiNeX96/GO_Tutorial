package main

import "fmt"

// Complete the function below to reverse a string

func ReverseString(s string) string {
    // Your code here
	// Tip : You can convert a string to a array of runes by doing []rune(s)
}

func main() {
    input := "hello"
    reversed := ReverseString(input)
    fmt.Println(reversed) // Should print "olleh"
}