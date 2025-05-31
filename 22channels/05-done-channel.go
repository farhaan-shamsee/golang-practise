package main // Declare the main package

import (
	"fmt" // Import the fmt package for formatted I/O
	"time"
)

func doWork(done <-chan bool)  {
	for {
		select {
		case <- done:
			return
		default:
			fmt.Println("DOING WORK")
		}
	}
	
}

func main() { // Main function

	doneChannel := make(chan bool)

	go doWork(doneChannel)

	time.Sleep(time.Second * 3)

	close(doneChannel)

}

// This Go program demonstrates how to use goroutines and channels to control the execution of a concurrent task.

// Step-by-step explanation:

// Channel Creation:
// doneChannel := make(chan bool)
// A channel named doneChannel is created. It is used to signal when the worker goroutine should stop.

// Starting a Goroutine:
// go doWork(doneChannel)
// The doWork function is started as a goroutine. This means it runs concurrently with the main function.

// Worker Function (doWork):

// It enters an infinite loop.
// Inside the loop, it uses a select statement:
// If a value is received from the done channel, the function returns and the goroutine stops.
// Otherwise (default), it prints "DOING WORK" and continues looping.
// Main Function Sleeps:
// time.Sleep(time.Second * 3)
// The main function waits for 3 seconds, allowing the goroutine to print "DOING WORK" repeatedly.

// Stopping the Goroutine:
// close(doneChannel)
// After 3 seconds, the main function closes the doneChannel. This causes the doWork goroutine to exit its loop and stop.

// Summary
// The program starts a background worker that prints "DOING WORK" repeatedly.
// After 3 seconds, the main function signals the worker to stop by closing the channel.
// The worker goroutine detects the signal and exits.