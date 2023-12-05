package main

import ("fmt"
		"time")

// Signals exemple

func main() {

	signal := make(chan string)

	go func () {
		fmt.Println(" GO routine Working...")
		time.Sleep(5 * time.Second)
		signal <- "Work done"
	}()

	variavel :=<- signal

	fmt.Println(variavel)

	fmt.Println("Routine ended!") 

	fmt.Println("Exiting") 

	


}