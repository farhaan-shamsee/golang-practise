package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	n := 36 // Original number to find all divisors
	divisors := []int{}

	// Loop from 1 to n to find all divisors of n
	// for i := 1; i <= n; i++ {
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ { 					//less number of steps
		// If n is divisible by i, then i is a divisor of n
		if n%i == 0 {
			divisors = append(divisors, i)
			if n/i != i {
				divisors = append(divisors, n/i)
			}
		}
	}
	slices.Sort(divisors)
	fmt.Println(divisors)
}

// Time complexity is O(N)
// This program prints all divisors of the number n
