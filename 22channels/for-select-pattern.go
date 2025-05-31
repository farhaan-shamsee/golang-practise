package main // Declare the main package

import (
	"fmt" // Import the fmt package for formatted I/O
)

func main() { // Main function

	charChannel := make(chan string, 3) // Create a buffered channel of strings with capacity 3

	chars := []string{"a", "b", "c"} // Define a slice of strings

	for _, char := range chars { // Iterate over each character in the slice
		select { // Use select statement (only one case)
		case charChannel <- char: // Send the character to the channel
		}
	}
	close(charChannel) // Close the channel after sending all characters

	for result := range charChannel { // Iterate over values received from the channel
		fmt.Println(result) // Print each value
	}

}
