package main

import (
	"fmt"
)

func main() {
	n := 5 // Original number to check if it is prime

	counter := 0 // Counter to count the number of divisors

	// Loop from 1 to sqrt(n) to find all divisors of n
	// For each i, if n is divisible by i, then both i and n/i are divisors
	for i := 1; i*i <= n; i++ {
		// If n is divisible by i, then i is a divisor of n
		if n%i == 0 {
			counter++ // Count divisor i
			if n/i != i {
				counter++ // Count divisor n/i if it's different from i
			}
		}
	}

	// If the number of divisors is exactly 2 (1 and n), then n is prime
	if counter == 2 {
		fmt.Println("Number is prime")
	} else {
		fmt.Println("Number is not prime")
	}
}

// Time complexity is O(sqrt(N))
// This program checks if n is prime by counting its divisors
// It does NOT print all divisors of n
