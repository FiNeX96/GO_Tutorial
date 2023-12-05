package main

import "fmt"

func RemoveDuplicates(nums []int) []int {
    unique := make(map[int]bool)
    result := []int{}

    for _, num := range nums {
        if !unique[num] {
            result = append(result, num)
            unique[num] = true
        }
    }

    return result
}

func main() {
    numbers := []int{1, 2, 2, 3, 4, 4, 5}
    result := RemoveDuplicates(numbers)
    fmt.Println(result) // Prints [1 2 3 4 5]
}
