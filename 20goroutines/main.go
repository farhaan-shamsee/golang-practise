package main

import (
	"fmt"
	"sync"
	// "log"
	"net/http"
	// "time"
)

// `signals` is a slice to store the endpoints that are successfully processed.
var signals = []string{"test"}

// `wg` is a WaitGroup to ensure the main function waits for all goroutines to complete.
var wg sync.WaitGroup // usually these are pointers

// `mut` is a Mutex to ensure thread-safe access to shared resources like `signals`.
var mut sync.Mutex // usually these are pointers

func main() {
	// List of websites to check their HTTP status codes.
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	// Loop through each website and start a goroutine to fetch its status code.
	for _, web := range websiteList {
		go getStatusCode(web) // Launch a goroutine for each website.
		wg.Add(1)             // Increment the WaitGroup counter for each goroutine.
	}

	// Wait for all goroutines to finish before exiting the main function.
	wg.Wait()
	fmt.Println(signals) // Print the final `signals` slice.
}

// `getStatusCode` fetches the HTTP status code for a given endpoint.
func getStatusCode(endpoint string) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes.

	// Perform an HTTP GET request to the endpoint.
	result, err := http.Get(endpoint)

	if err != nil {
		// Print an error message if the request fails.
		fmt.Printf("OOPS in endpoint: %s\n", endpoint)
	} else {
		// Lock the mutex to safely update the shared `signals` slice.
		mut.Lock()
		signals = append(signals, endpoint) // Append the endpoint to the `signals` slice.
		mut.Unlock()                        // Unlock the mutex after updating the slice.

		// Print the HTTP status code for the endpoint.
		fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
	}
}
