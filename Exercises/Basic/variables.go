package main 

import "fmt"

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

	// var testing := 5  this wont work, the compiler will ask you to use '=' instead of ':='
	// int testing = 5 this wont work either, you cant use the keyword 'int' to declare a variable
	

	fmt.Println(x,y,hello,world,pi,pi2, test)

}