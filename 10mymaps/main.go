package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang!")

	// Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).
	// A map is an unordered collection of key/value pairs, where each key is unique.

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("list of all languages: ", languages)
	fmt.Println("JS shorts for", languages["JS"])

	delete(languages, "RB")
	fmt.Println("list of all languages after deleting RB: ", languages)

	// looping through maps

	for key, value := range languages { 
		fmt.Printf("For key %v, value is %v\n", key, value); // prints key and value
	}
	
	// if you want to print only values, you can do it like this
	// := is called walrus operator
	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value); // prints only value
	}

	// To remove all key/value pairs from a map, use the clear builtin.
	// clear(languages)
}