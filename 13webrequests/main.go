package main

import (
	"fmt"
	"io"
	"net/http"
)

// Define a constant for the URL to make a GET request
const url = "https://google.com"

func main() {
	// Print a welcome message
	fmt.Println("Welcome to webrequest")

	// Make an HTTP GET request to the specified URL
	response, err := http.Get(url)
	if err != nil {
		// If there's an error during the GET request, terminate the program
		panic(err)
	}

	// Print the type of the response object
	fmt.Printf("Response is of type: %T\n", response)
	// Print the response object itself (contains metadata about the HTTP response)
	fmt.Println("Response is: ", response)

	// Ensure the response body is closed after the function exits
	defer response.Body.Close()

	// Read the response body into a byte slice
	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		// If there's an error reading the response body, terminate the program
		panic(err)
	}

	// Convert the byte slice to a string and print the content of the response body
	fmt.Println("content is: ", string(databytes))
}