package main 


import ("fmt")


/*

Maps are a data structure that store key-value pairs, just like you know from other languages.


The syntax to create a map is :

map_name := make(map[key_type]value_type) 

This declares and initializes the map with the given key and value types.

To add a key-value pair to a map, you can do the following:

map_name[key] = value

To delete a value from a map, you can do the following:

delete(map_name, key)

You can check if a key exists in a map by doing the following:

value, exists := map_name[key]

if exists {
	// the key exists
}

*/


func main() {


	// Start by creating a map with key type string and value type int
	ages := make(map[string]int)

	// Now add some key-value pairs to the map ( you can assume its people and their age )
	ages["Alice"] = 15
	ages["Bob"] = 40

	// Print the map
	for name, age := range ages {
		fmt.Printf("%s: %d\n", name, age)
	}

	// Now delete a key-value pair from the map
	delete(ages, "Bob")

	// Print the map again
	for name, age := range ages {
		fmt.Printf("%s: %d\n", name, age)
	}



}