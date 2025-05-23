package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Print welcome message
	fmt.Println("Welcome to our Pizza Application")
	// Ask user to rate the pizza
	fmt.Println("Please rate our pizza between 1 and 5")

	// Create a new reader to read input from standard input (keyboard)
	reader := bufio.NewReader(os.Stdin)

	// Read the user's input until a newline character is encountered
	input, _ := reader.ReadString('\n')

	// Print a thank you message along with the user's input
	fmt.Println("Thanks for rating us", input)

	// Convert the input string (after trimming spaces/newline) to a float64
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// Check if there was an error during conversion
	if err != nil {
		// Print the error if conversion failed
		fmt.Println(err)
	} else {
		// If conversion succeeded, add 1 to the rating and print the result
		fmt.Println("Added 1 to your rating", numRating+1)
	}
}