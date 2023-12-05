package main

import ("fmt"
		"time")

// Signals exemple

func main() {

	signal := make(chan string)

	go func () {
		fmt.Println(" GO routine Working...")
		time.Sleep(20 * time.Second)
		signal <- "Work done"
	}()

	<- signal

	fmt.Println("Routine ended!") 

	fmt.Println("Exiting") 

	


}