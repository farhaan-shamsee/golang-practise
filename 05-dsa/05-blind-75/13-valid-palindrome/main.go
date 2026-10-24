/*
Problem: Valid Palindrome
A phrase is a palindrome if, after converting all uppercase letters to lowercase and
removing all non-alphanumeric characters, it reads the same forward and backward.

Given a string s, return true if it is a palindrome, or false otherwise.

Real-world use case for Platform Engineers:
- Data integrity checks in configuration validation
- String manipulation efficiency in log processing
- Verifying symmetric patterns in network configurations
- Validating reversible token formats in authentication

Approaches:

1. Two Pointers (Optimal):
  - Filter and normalize string
  - Use two pointers from both ends
  - Time: O(n), Space: O(1) if in-place, O(n) if creating new string

2. Reverse and Compare:
  - Create cleaned version of string
  - Reverse it and compare
  - Time: O(n), Space: O(n)

Algorithm (Two Pointers):
 1. Initialize left = 0, right = len(s) - 1
 2. While left < right:
    a. Skip non-alphanumeric from left
    b. Skip non-alphanumeric from right
    c. Compare characters (case-insensitive)
    d. If different, return false
    e. Move pointers inward
 3. Return true if all matched

Time Complexity: O(n) - single pass with two pointers
Space Complexity: O(1) - only using pointers

Example 1: s = "A man, a plan, a canal: Panama"

Step 1 - Normalize (conceptually):

	Original: "A man, a plan, a canal: Panama"
	Cleaned:  "amanaplanacanalpanama"

Step 2 - Two pointer comparison:

	Position: 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20
	String:   a  m  a  n  a  p  l  a  n  a  c  a  n  a  l  p  a  n  a  m  a
	Left:     ↑                                                            ↑  Right
	          a == a ✓

	Position: 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20
	String:   a  m  a  n  a  p  l  a  n  a  c  a  n  a  l  p  a  n  a  m  a
	Left:        ↑                                                      ↑     Right
	             m == m ✓

	... (continue) ...

	All characters match → return true

Example 2: s = "race a car"

	Cleaned: "raceacar"

	Position: 0  1  2  3  4  5  6  7
	String:   r  a  c  e  a  c  a  r
	Left:     ↑                    ↑  Right
	          r == r ✓

	Position: 0  1  2  3  4  5  6  7
	String:   r  a  c  e  a  c  a  r
	Left:        ↑              ↑     Right
	             a == a ✓

	Position: 0  1  2  3  4  5  6  7
	String:   r  a  c  e  a  c  a  r
	Left:           ↑        ↑        Right
	                c == c ✓

	Position: 0  1  2  3  4  5  6  7
	String:   r  a  c  e  a  c  a  r
	Left:              ↑  ↑           Right
	                   e != a ✗

	Mismatch found → return false
*/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	// Convert to lowercase
	s = strings.ToLower(s)

	left := 0
	right := len(s) - 1

	for left < right {
		// Skip non-alphanumeric from left
		for left < right && !isAlphanumeric(rune(s[left])) {
			left++
		}

		// Skip non-alphanumeric from right
		for left < right && !isAlphanumeric(rune(s[right])) {
			right--
		}

		// Compare characters
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphanumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func main() {
	test1 := "A man, a plan, a canal: Panama"
	fmt.Printf("Input: \"%s\"\n", test1)
	fmt.Printf("Output: %v\n\n", isPalindrome(test1)) // true

	test2 := "race a car"
	fmt.Printf("Input: \"%s\"\n", test2)
	fmt.Printf("Output: %v\n\n", isPalindrome(test2)) // false

	test3 := " "
	fmt.Printf("Input: \"%s\"\n", test3)
	fmt.Printf("Output: %v\n\n", isPalindrome(test3)) // true

	test4 := "Was it a car or a cat I saw?"
	fmt.Printf("Input: \"%s\"\n", test4)
	fmt.Printf("Output: %v\n\n", isPalindrome(test4)) // true

	// Real-world example: Symmetric configuration key
	configKey := "api-key-ipa"
	fmt.Printf("Config key: \"%s\"\n", configKey)
	fmt.Printf("Is palindrome: %v\n", isPalindrome(configKey)) // true
}
