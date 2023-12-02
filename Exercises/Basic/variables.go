package main 

import ("fmt")

func main() {

	/* 
	   in Go you can declare variables in a few ways
	   You can declare a variable with a type and a value
	   or you can just declare it with a value and let the compiler figure out the type
	*/

	// declare a variable with a type and a value

	var x int = 5


	// declare a variable with a value and let the compiler figure out the type

	y := 10  // you can use := and drop the keyword "var" all together

	// The usual data types are available in Go aswell, here are some examples

	hello := true  
	var world = "Hello World"
	pi := float32(3.14)
	var pi2 = float64(3.14)
	var test string = "test"

	/* 
	
	var testing := 5  this wont work, the compiler will ask you to use '=' instead of ':='
	int testing = 5 this wont work either, you cant use the keyword 'int' to declare a variable
	testing := 5 this will work, the compiler will figure out the type of the variable
	var testing = 5 will work aswell

	*/

	var not_a_char rune = 'a' 
	
	/* rune is an alias for int32 and is used to represent a Unicode code point

	   Runes are basically 'char' from other languages, but with differences in encoding and representation

	*/

	// With runes we can write some weird stuff like this

	const string1 = "สวัสดีцика"



	for idx, runeValue := range string1 {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

	
	fmt.Println(x,y,hello,world,pi,pi2, test , not_a_char)

}