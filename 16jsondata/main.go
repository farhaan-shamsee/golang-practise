package main

import (
	"encoding/json"
	"fmt"
)

// Define a struct to represent a course with JSON tags for serialization
type course struct {
	Name     string   `json:"coursename"` // Maps the Name field to "coursename" in JSON
	Price    int      // Price will be serialized with the same name
	Platform string   `json:"website"`   // Maps the Platform field to "website" in JSON
	Password string   `json:"-"`         // Password will be ignored during JSON serialization
	Tags     []string `json:"tags,omitempty"` // Tags will be omitted if nil or empty
}

func main() {
	// Entry point of the program
	fmt.Println("Welcome to JSON data!")
	EncodeJson() // Call the function to encode data into JSON
}

func EncodeJson() {
	// Create a slice of course structs with sample data
	mycourse := []course{
		{"ReactJS Bootcamp", 299, "Udemy", "abc123", []string{"web dev", "js"}},
		{"MERN Stack", 199, "Udemy", "xyz456", []string{"web dev", "js"}},
		{"Angular Bootcamp", 299, "Udemy", "def789", nil}, // Tags are nil here
	}

	// Convert the slice of courses to JSON with indentation for readability
	// json.MarshalIndent adds indentation to the JSON output
	finalJson, err := json.MarshalIndent(mycourse, "", "\t")

	// Handle any error that occurs during JSON encoding
	if err != nil {
		panic(err) // Terminate the program if an error occurs
	}

	// Print the JSON output as a string
	fmt.Printf("Body: %s\n", finalJson)
}
func DecodeJson() {
	// Sample JSON data coming from the web
	jsonData := []byte(`
	[
		{
			"coursename": "ReactJS Bootcamp",
			"Price": 299,
			"website": "Udemy",
			"tags": ["web dev", "js"]
		},
		{
			"coursename": "MERN Stack",
			"Price": 199,
			"website": "Udemy",
			"tags": ["web dev", "js"]
		},
		{
			"coursename": "Angular Bootcamp",
			"Price": 299,
			"website": "Udemy"
		}
	]`)

	// Define a variable to hold the decoded data
	var courses []course

	// Decode the JSON data into the courses slice
	err := json.Unmarshal(jsonData, &courses)
	if err != nil {
		panic(err) // Terminate the program if an error occurs
	}

	// Print the decoded data
	for _, course := range courses {
		fmt.Printf("Course: %+v\n", course)
	}
}