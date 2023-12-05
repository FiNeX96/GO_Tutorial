package main

import "fmt"

/* The syntax for functions is as follows:

	func name_of_function(parameters) (return type, error) {
		// code
	}


	For object methods:

	func (object_name object_type) name_of_function(parameters) (return type, error) {
		// code
	}

*/



// start by declaring a function to add two numbers


func add(x1,x2 int ) (int){
	return x1+x2
}



// now declare a function that takes a string as input and returns the length of the string
// hint: consider the python way to calculate the length =)


func strlen(s1 string)(int){
	return len(s1)
}




// Use this main function to test your functions
func main() {
	

	fmt.Println("Hello World") // you can delete this line after u insert some code here 

	fmt.Println(add(1,2))
	
	str := "ola"

	value := strlen(str)

	fmt.Println(value)

	
}