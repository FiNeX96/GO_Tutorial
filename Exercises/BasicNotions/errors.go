package main

import (
	"errors"
	"fmt"
)

// Here we have a simple divide function
// As you can see, there is a error dataype which we can use to return errors
func divide(x, y int) (int, error) {
	if y == 0 {
		// Return an error if the divisor is zero
		return 0, errors.New("division by zero")
	}
	result := x / y
	return result, nil

	/*
	In Go, nil is a predeclared identifier that represents a "zero value" or "null" value for pointers, slices, maps, channels, and function types.
	Its frequently used to represent a absence of a value ( null ) or to handle errors
	*/
}

func main() {
	// Example 1: Successful division
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Example 2: Division by zero (error case)
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	/*

	Most GO libraries implement this structure for error handling, so you can use this to handle errors in your code

	result, err = operation (...)

	For example, the os.Open function returns a file and an error:

	file, err := os.Open("filename.ext")

	if err != nil {
		// handle the error here
	}

	*/
}
