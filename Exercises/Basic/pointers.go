package main 


import ("fmt")

// Go supports pointers, but not pointer arithmetic
// You can use '*' or '&' to get the value of a pointer or its address respectively, like in c(++)

func main() {


	p := 5 // declare variable p with value 5

	fmt.Println(p) // print variable p

	fmt.Println(&p) // print the address of variable p

	ptr := *(&p) // pointer to variable p

	fmt.Println(ptr) // print the value of the pointer ptr

	fmt.Println(&ptr) // print the address of the pointer ptr

}