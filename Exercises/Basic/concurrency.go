package main 


import ("fmt"
		"sync"
		"time"
		"math/rand"
	)


/*

GO has a built-in concurrency model called goroutines and channels.

A goroutine is a lightweight thread managed by the Go runtime.

A channel is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine.

The syntax to create a goroutine is:

go function_name(parameters)

However, the main function ( main routine ) does not wait for the goroutine to finish, 
so we need to use a waitgroup to make sure the main routine waits for the goroutine to finish.


Analyse the functions below and try to understand what they do.

*/


func main() {

	// Create a waitgroup, that is used to wait for all goroutines to finish before exiting the main function/routine
	var wg sync.WaitGroup

	counter := 1

	wg.Add(2) // This wait group will wait for 2 goroutines to finish before exiting the main function/routine

	// Start the first goroutine
	go dowork(&wg, counter)

	counter++


	// Start the second goroutine
	go dowork(&wg, counter)


	// Tell the main routine to wait for all this groups goroutines to finish
	wg.Wait()

	fmt.Println("All go routines finished !")
}


func dowork(wg *sync.WaitGroup , counter int) {

	fmt.Println("Go routine " , counter ," gonna do some work !")
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Println("Go routine, " , counter , "finished !")
	wg.Done()
}