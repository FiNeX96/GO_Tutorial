package main 


import (
		"fmt"
     	"io/ioutil" 
		"os"
		 ) 


func main() {


	// print a basic string to the terminal

	fmt.Println("Hello World")	


	// take as output a string/int and print it to the terminal

	var x int 
	fmt.Print("Enter a number: ")
	fmt.Scan(&x)
	fmt.Println("Your number -> " ,x)


	// read data from a file and print it to a terminal
	// you may use the file test.txt

	data, err := ioutil.ReadFile("test.txt")

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fmt.Println("Contents of file:", string(data))


	// write to a file
	
	file, err := os.OpenFile("test2.txt", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	file.Write([]byte("appended some data\n"))
	


}