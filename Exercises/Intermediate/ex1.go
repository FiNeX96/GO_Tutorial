package main

import "fmt"



// Complete the function below to reverse a string

func ReverseString(s string) string {
    // Your code here
	// Tip : lookup the rune datatype -> https://www.geeksforgeeks.org/rune-in-golang/
}

func main() {
    input := "hello"
    reversed := ReverseString(input)
    fmt.Println(reversed) // Should print "olleh"
}