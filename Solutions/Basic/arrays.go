package main

import (
	"fmt"
)

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	// Declare an array with size 5, type int, and give it some starting values
	arr := [5]int{1, 2, 3, 4}

	// Print the array
	fmt.Println("Original Array:", arr)

	// Convert array to a slice to be able to append to it
	arr_sliced := arr[:]

	// Append a value to the slice
	arr_sliced = append(arr_sliced, 6)

	// Print the slice after appending
	fmt.Println("Array after appending 6:", arr_sliced)

	// Create a function to remove a specific index from the array
	indexToRemove := 2
	arr_sliced = removeIndex(arr_sliced, indexToRemove)

	// Print the array with the removed index
	fmt.Printf("Array after removing index %d: %v\n", indexToRemove, arr_sliced)
}
