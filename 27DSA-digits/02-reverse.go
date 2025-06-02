package main

import "fmt"

func main() {
	n := 778090           // Original number to reverse
	revNumber := 0        // Variable to store the reversed number

	// Loop until all digits are processed
	for n > 0 {
		lastDigit := n % 10           // Extract the last digit
		n = n / 10                    // Remove the last digit from n
		revNumber = revNumber*10 + lastDigit // Append lastDigit to revNumber
	}

	fmt.Println("Reverse number:", revNumber) // Output the reversed number
}
