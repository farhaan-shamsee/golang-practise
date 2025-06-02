package main

import "fmt"
import "math"

// main is the entry point of the program
func main() {
	n := 2397 // TODO: Try with different values, including negatives and zero

	counter := 0
	if n == 0 {
		counter = 1 // Special case: 0 has 1 digit
	} else {

		// TODO: Handle negative numbers (currently only works for positive n)
		for n > 0 {
			n = n / 10            // remove the last digit
			counter = counter + 1 // increment digit count
		}
	}
	fmt.Println("digits: ", counter) // print total digits

	// TODO: Handle negative numbers (currently not handled)
	// TODO: Refactor digit counting logic into a reusable function
	// TODO: Add unit tests for various cases (positive, negative, zero, large numbers)
	
	// Alternative approach
	count := int(math.Log10(float64(n))) + 1
	fmt.Println("digits: ", count)

	// Notes for revision:
	// - The number of iterations depends on how many times n can be divided by 10 before reaching 0.
	// - This means the time complexity is O(log10(N)).
	// - If dividing by 2, time complexity would be O(log2(N)).
	// - Remember to handle edge cases (negative numbers, zero).
}
