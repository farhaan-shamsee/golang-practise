package main

import "fmt"
import "math"

func main() {
	n := 1634 // Original number to check for Armstrong property
	sum := 0
	duplicateN := n // Store the original number for comparison later

	digits := int(math.Log10(float64(n))) + 1 // Calculate the number of digits in n

	// Loop to process each digit of the number
	for n > 0 {
		lastDigit := n % 10 // Extract the last digit
		n = n / 10          // Remove the last digit from n
		// Add the digit raised to the power of 'digits' to sum
		sum += int(math.Pow(float64(lastDigit), float64(digits)))
	}

	// Check if the sum equals the original number (Armstrong condition)
	if sum == duplicateN {
		fmt.Println("number is armstrong, sum is: ", sum)
	} else {
		fmt.Println("Number is not armstrong, sum is: ", sum)
	}
}
