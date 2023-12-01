package main

import (
	"fmt"
)

func main() {
	// Variables and Constants
	var a, b int = 5, 3
	const pi = 3.14159

	// Arithmetic Operations, just like the ones in other languages
	sum := a + b
	difference := a - b
	product := a * b
	quotient := float64(a) / float64(b) // Convert to float64 for precise division

	// Print Results
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)

	// Conditional Statements
	if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("b is greater than or equal to a")
	}

	// Loops
	fmt.Println("Numbers from 1 to 5:")
	for   i := 1; i <= 5; i++ { // may also use for 
		fmt.Print(i, " ")
	}
	fmt.Println()


}


