package main

import "fmt"

// Define a struct named 'Person'
type Person struct {
    Name    string
    Age     int
    Address string
}

// Define a method 'PrintDetails' for the 'Person' struct
func (p Person) PrintDetails() {
    fmt.Printf("Name: %s\nAge: %d\nAddress: %s\n", p.Name, p.Age, p.Address)
}

func main() {
    // Create an object of the 'Person' struct
    person1 := Person{
        Name:    "John Doe",
        Age:     30,
        Address: "123 Main St",
    }

    // Print the attributes of the object
    fmt.Println("Attributes of the person1 object:")
    fmt.Println("Name:", person1.Name)
    fmt.Println("Age:", person1.Age)
    fmt.Println("Address:", person1.Address)

    // Call the method 'PrintDetails' on the object
    fmt.Println("\nCalling PrintDetails method:")
    person1.PrintDetails()
}
