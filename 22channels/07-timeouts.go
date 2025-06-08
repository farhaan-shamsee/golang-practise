package main

import (
	"fmt"
	"time"
)

func main() {

	// Create a buffered channel c1
	c1 := make(chan string, 1)
	// Start a goroutine that sleeps for 2 seconds, then sends a value to c1
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Use select to wait for either a value from c1 or a 1-second timeout
	select {
	case res := <-c1:
		fmt.Println(res) // Prints the result if received before timeout
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1") // Prints timeout if no result in 1 second
	}

	// select waits for communication and handles multiple channel operations.
	// case res := <-c1: Tries to receive a value from the channel c1. Prints the received value.
	// case <-time.After(1 * time.Second): Waits for 1 second. If no message is received from c1 within this time, it times out and "timeout 1" is printed.
	// Outcome: Since the goroutine sleeps for 2 seconds, the first select times out, so "timeout 1" is printed.

	
	// ----------------------------------------------------------------------------------------------------------------------------
	// ----------------------------------------------------------------------------------------------------------------------------
	// ----------------------------------------------------------------------------------------------------------------------------


	// Create a buffered channel c2
	c2 := make(chan string, 1)
	// Start a goroutine that sleeps for 2 seconds, then sends a value to c2
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	
	// Use select to wait for either a value from c2 or a 3-second timeout
	select {
	case res := <-c2:
		fmt.Println(res) // Prints the result if received before timeout
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2") // Prints timeout if no result in 3 seconds
	}
	
	// 	Works the same way as the first but with a 3-second timeout.
	// Outcome: The goroutine finishes after 2 seconds, so "result 2" is received from c2, and "result 2" is printed before the 3-second timeout is reached.
	
}