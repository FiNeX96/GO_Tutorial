package main

import ("fmt")

/* The syntax to define an array is:

var array_name [size] array_type

- var arr [5] int


You can also define multidimensional arrays, like this:

- var arr2 [5][5] int

Go uses the following definitions:

- array : fixed length  -> [5]int{1,2,3,4,5}

- slice : dynamic length -> []int{1,2,3,4,5}

If you have a array, you can easily create a slice of it :

arr := [5]int{1,2,3,4,5}

slice := arr[:] for all elements

slice := arr[1:3] for elements index 1 and 2

Some operations can only be done on slices, such as appending.

The syntax to append to a slice is:

slice = append(slice, element) - you can only append to slices, not arrays !

You can also append slices to other slices, with the following syntax:

slice = append(slice, slice2...) - this tells the compiler to unpack slice2 and append each element to slice


*/

func main(){


	// start by declaring a array with size 5 and type int and give it some starting values


	// print the array


	// append a value to the array ( remember, you need to convert it to a slice first ! )


	// now create a function to remove a specific index from the array ( tip : use slices and append )
	// This one may be a bit tricky, feel free to ask for help =)


	// print the array/slice with the removed index again




}
