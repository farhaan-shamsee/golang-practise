package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// Entry point of the program
	fmt.Println("Welcome to webrequest!")
}

func PerformGetRequest() {
	// Define the URL to send the GET request to
	const myurl = "http://localhost:8000/get"
	
	// Perform the GET request
	response, err := http.Get(myurl)
	if err != nil {
		// Handle any errors that occur during the GET request
		panic(err)
	}
	// Ensure the response body is closed after the function completes
	defer response.Body.Close()

	// Print the HTTP status code of the response
	fmt.Println("Status code: ", response.StatusCode)
	// Print the content length of the response
	fmt.Println("Content length: ", response.ContentLength)

	// Create a strings.Builder to build the response string
	var responseString strings.Builder

	// Read the response body into a byte slice
	content, _ := io.ReadAll(response.Body)

	// Method 1: using string(content)
	// Uncomment these lines to print the raw content and its string representation
	// fmt.Println(content)
	// fmt.Println(string(content))
	
	// Method 2: using strings.Builder
	// Write the content to the strings.Builder and get the byte count
	byteCount, _ := responseString.Write(content)
	// Print the number of bytes written to the strings.Builder
	fmt.Println("Byte count is: ", byteCount)
	// Print the final response string
	fmt.Println("Response string is: ", responseString.String())
}