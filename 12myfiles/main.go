package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Print a welcome message
	fmt.Println("Welcome to files in go")

	// Define the content to be written to the file
	content := "This will be printed in the file"

	// Create a new file named "myfile.txt"
	file, err := os.Create("./myfile.txt")

	// Check if there was an error while creating the file
	if err != nil {
		panic(err)
	}

	// Write the content to the file
	length, err := io.WriteString(file, content)

	// Check for any error during the write operation
	checkNilErr(err)

	// Print the length of the content written to the file
	fmt.Println("Length is: ", length)

	// Close the file using defer to ensure it is closed after the function exits
	defer file.Close()

	// Call the reader function to read the content of another file
	reader("./myfile.txt")
}

// Function to read the content of a file
func reader(filename string) {
	// Read the entire content of the file into a byte slice
	databyte, err := os.ReadFile(filename)

	// Check for any error during the file read operation
	checkNilErr(err)

	// Print the content of the file as a string
	fmt.Println("text data is: \n", string(databyte))
}

// Function to check if an error is nil, and panic if it is not
func checkNilErr(Error error) {
	if Error != nil {
		panic(Error)
	}
}