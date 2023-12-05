package main

import "fmt"

// Complete the function below to reverse a string

func ReverseString(s string) string {
    // Your code here
	// Tip : You can convert a string to a array of runes by doing []rune(s)
    runes := [] rune(s)

    for i,j := 0, len(runes)-1;i<j; i,j = i+1, j-1{
        runes[i],runes[j] = runes[j],runes[i]
    }

    return string(runes)
}

func main() {
    input := "hello"
    reversed := ReverseString(input)
    fmt.Println(reversed) // Should print "olleh"
}