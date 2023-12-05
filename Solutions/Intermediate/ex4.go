package main

import "fmt"

func Factorial(n int) int {
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}

func main() {
    num := 5
    result := Factorial(num)
    fmt.Println(result) // Prints 120
}
