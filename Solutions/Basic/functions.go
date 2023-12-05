package main

import "fmt"

// Function to add two numbers
func add(a, b int) int {
    return a + b
}

// Function to calculate the length of a string
func stringLength(s string) int {
    return len(s)
}

// Use this main function to test your functions
func main() {
    // Test the add function
    result := add(3, 5)
    fmt.Println("Result of adding 3 and 5:", result)

    // Test the stringLength function
    str := "Hello, Gophers!"
    length := stringLength(str)
    fmt.Printf("Length of the string '%s': %d\n", str, length)
}
