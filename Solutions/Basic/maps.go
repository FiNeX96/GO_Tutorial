package main

import "fmt"

func main() {
    // Create a map with key type string and value type int
    ages := make(map[string]int)

    // Add some key-value pairs to the map (people and their age)
    ages["Alice"] = 30
    ages["Bob"] = 25
    ages["Charlie"] = 35

    // Print the map
    fmt.Println("Map after adding key-value pairs:")
    for name, age := range ages {
        fmt.Printf("%s: %d\n", name, age)
    }

    // Delete a key-value pair from the map
    delete(ages, "Bob")

    // Print the map again
    fmt.Println("\nMap after deleting the key 'Bob':")
    for name, age := range ages {
        fmt.Printf("%s: %d\n", name, age)
    }
}
