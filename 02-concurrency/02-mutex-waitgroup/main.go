package main

import (
	"fmt"  // Importing the fmt package for formatted I/O
	"sync" // Importing the sync package for synchronization primitives
)

func main() {
	fmt.Println("Race condition") // Print a message indicating the topic

	wg := &sync.WaitGroup{} // Create a WaitGroup to wait for all goroutines to finish
	mut := &sync.Mutex{}    // Create a Mutex to prevent race conditions

	var score = []int{0} // Initialize a slice to store scores

	wg.Add(3) // Add 3 to the WaitGroup counter, as we will launch 3 goroutines

	// First goroutine
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")     // Print a message indicating the first goroutine
		mut.Lock()               // Lock the mutex to ensure exclusive access to the score slice
		score = append(score, 1) // Append 1 to the score slice
		mut.Unlock()             // Unlock the mutex
		wg.Done()                // Decrement the WaitGroup counter
	}(wg, mut) // Pass the WaitGroup and Mutex as arguments to the goroutine

	// Second goroutine
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")     // Print a message indicating the second goroutine
		mut.Lock()               // Lock the mutex to ensure exclusive access to the score slice
		score = append(score, 2) // Append 2 to the score slice
		mut.Unlock()             // Unlock the mutex
		wg.Done()                // Decrement the WaitGroup counter
	}(wg, mut) // Pass the WaitGroup and Mutex as arguments to the goroutine

	// Third goroutine
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")   // Print a message indicating the third goroutine
		mut.Lock()               // Lock the mutex to ensure exclusive access to the score slice
		score = append(score, 3) // Append 3 to the score slice
		mut.Unlock()             // Unlock the mutex
		wg.Done()                // Decrement the WaitGroup counter
	}(wg, mut) // Pass the WaitGroup and Mutex as arguments to the goroutine

	wg.Wait()          // Wait for all goroutines to finish
	fmt.Println(score) // Print the final score slice
}
