package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	// Create a buffered channel with capacity 2
	myCh := make(chan int, 2)

	// Uncommenting the below lines will cause a deadlock error because no goroutine is receiving from the channel.
	// Channels in Go require at least one sender and one receiver to avoid deadlocks.
	// myCh <- 5
	// fmt.Println(<-myCh)

	wg := &sync.WaitGroup{}

	// Add 2 to the WaitGroup counter to wait for two goroutines to complete
	wg.Add(2)

	// Goroutine to receive data from the channel (receive-only channel)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// Close the channel to indicate no more values will be sent
		close(myCh)

		// Receive a value from the channel and check if the channel is open
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen) // Prints false because the channel is closed
		fmt.Println(val)           // Prints 0 because the channel is closed and no value is available

		// Attempt to receive another value from the channel
		fmt.Println(<-myCh) // Prints 0 because the channel is closed

		// Mark this goroutine as done
		wg.Done()
	}(myCh, wg)

	// Goroutine to send data to the channel (send-only channel)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// Closing the channel here will cause a panic because you cannot send to a closed channel
		// close(myCh)

		// Uncommenting the below lines will send values to the channel
		// However, since the channel is closed above, these lines will cause a panic
		// myCh <- 0
		// myCh <- 5
		// myCh <- 6

		// Mark this goroutine as done
		wg.Done()
	}(myCh, wg)

	// Wait for both goroutines to complete
	wg.Wait()
}
