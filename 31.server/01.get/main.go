/*
"Write a function that takes a URL string and returns a specific query parameter value.

Input: 
  URL: "https://api.scout24.com/clusters?region=eu-west-1&status=active"
  Key: "region"

Output: "eu-west-1"

Handle: Missing key → return empty string

*/

package main

import (
	"fmt"
	"net/url"
)

func urlHandler(urlToParse string, keyToSearch string) string {
	// test := make(map[string]string)
	u, err := url.Parse(urlToParse)
	if err != nil {
		panic(err)
	}
	return u.Query().Get(keyToSearch)

}

func main() {
	response := urlHandler("https://api.scout24.com/clusters?region=eu-west-2&status=active", "region")
	fmt.Println(response)
}