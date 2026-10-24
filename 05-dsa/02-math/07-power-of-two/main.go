package main

import "fmt"

func main() {
	n := 64 // Number to check

	// Approach 1: using a loop
	// Keep dividing by 2 while n is even; a power of 2 ends at exactly 1
	isPowerLoop := false
	if n > 0 {
		temp := n
		for temp%2 == 0 {
			temp = temp / 2
		}
		isPowerLoop = temp == 1
	}

	// Approach 2: without a loop
	// A power of 2 has exactly one bit set, so n & (n-1) clears it to 0
	isPowerBitwise := n > 0 && (n&(n-1)) == 0

	fmt.Println("Power of 2 (loop):", isPowerLoop)
	fmt.Println("Power of 2 (bitwise):", isPowerBitwise)
}

// Loop approach: Time complexity O(log N)
// Bitwise approach: Time complexity O(1)
