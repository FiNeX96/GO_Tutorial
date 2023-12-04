package main

import ("fmt")


/*

	Go is not a pure object oriented language, but it does support objects using Structs like C
	A struct may be defined as follows:

	type struct_name struct {
		field1 type_field1
		field2 type_field2
		method1() type_method1
	}

	You can also nest structs inside structs, like this:


	type room struct {
		room_number int
	}

	type home struct {
		address string
		rooms room 
	}

	Interfaces also exist in go, they are defined as follows:

	type test interface {
		attribute1 type1
		method2() type2
		method3() type3
	}

	Now you can declare these interface methods

	func (test) method2() type2 {
		// code
	}

*/


// For the following exercises you may use the main function to test your code
// Start by creating a struct, with whatever name/attributes you want


// Now using the main function, create an object of the struct you just created


// Print the attributes of the object you just created, you can use the following syntax to get a attribute value -> object.attribute


// Now write a method upon the struct you created, the syntax is the one explained above ( for example, print all the attributes of the object )


// Finally, call the method you just created and print its result



func main() {


	fmt.Println("Hello World") // you can delete this line after u insert some code here



}