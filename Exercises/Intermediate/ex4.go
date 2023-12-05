package main

import "fmt"

// Complete the function below to calculate the factorial of a number

func Factorial(n int) int {
    // Your code here
    result :=1
    for i := 1; i<=n; i++{
        result *= i
    }
    return result
}

func main() {
    num := 5
    result := Factorial(num)
    fmt.Println(result) // Should print 120
}