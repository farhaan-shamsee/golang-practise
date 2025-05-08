package main

import (
	"fmt"
	"net/url"
)

// Define a constant URL string
const myurl string = "https://lco.dev:300/learn?coursename=reactjs&paymentid=ghb32adw"

func main() {
	fmt.Println("Welcome to URL handling!")
	fmt.Println("url is: ", myurl)

	// Parsing the URL
	result, _ := url.Parse(myurl)

	// Print the scheme (protocol) of the URL
	fmt.Println(result.Scheme)

	// Print the host (domain + port) of the URL
	fmt.Println(result.Host)

	// Print the path of the URL
	fmt.Println(result.Path)

	// Print the port of the URL
	fmt.Println(result.Port())

	// Print the raw query string of the URL
	fmt.Println(result.RawQuery)

	// Extract query parameters as a map
	qparams := result.Query()
	fmt.Printf("the type of query params are: %T\n", qparams)

	// Access specific query parameters by their keys
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	// Iterate over all query parameters and print their values
	for _, val := range qparams {
		fmt.Println("Value of params is: ", val)
	}

	// Construct a new URL from its parts
	partsOfUrl := &url.URL{
		Scheme:  "https",          // Protocol
		Host:    "lco.dev",        // Domain
		Path:    "/tutcss",        // Path
		RawPath: "user=farrow",    // Raw path (not commonly used)
	}

	// Convert the URL struct to a string
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}