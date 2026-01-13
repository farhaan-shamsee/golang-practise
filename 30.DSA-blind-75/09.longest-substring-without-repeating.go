/*
Problem: Longest Substring Without Repeating Characters
Given a string s, find the length of the longest substring without repeating characters.

Real-world use case for Platform Engineers:
- Handling continuous data streams in developer portals
- Analyzing unique event sequences in observability tools
- Processing unique request patterns in API gateways
- Finding longest unique token sequences in authentication flows

Approaches:

1. Brute Force:
  - Check all substrings and verify uniqueness
  - Time: O(n³), Space: O(min(n, m)) where m = charset size
  - Not practical

2. Sliding Window with Hash Set (Optimal):
  - Use two pointers (left, right) to form a window
  - Expand window by moving right pointer
  - If duplicate found, shrink from left until duplicate removed
  - Track maximum window size
  - Time: O(n), Space: O(min(n, m))

3. Sliding Window with Hash Map (Optimized):
  - Store character and its index
  - When duplicate found, jump left pointer directly
  - Time: O(n), Space: O(min(n, m))

Algorithm (Sliding Window with Hash Set):
 1. Initialize left = 0, maxLen = 0, and empty set
 2. For right from 0 to n-1:
    a. While s[right] exists in set:
    - Remove s[left] from set
    - Increment left
    b. Add s[right] to set
    c. Update maxLen = max(maxLen, right - left + 1)
 3. Return maxLen

Time Complexity: O(n) - each character visited at most twice
Space Complexity: O(min(n, m)) - m is charset size (26 for lowercase letters)

Example: s = "abcabcbb"

Initial: left=0, maxLen=0, set={}

right=0: char='a'

	'a' not in set → add 'a' → set={'a'}
	window = "a", len=1, maxLen=1

right=1: char='b'

	'b' not in set → add 'b' → set={'a','b'}
	window = "ab", len=2, maxLen=2

right=2: char='c'

	'c' not in set → add 'c' → set={'a','b','c'}
	window = "abc", len=3, maxLen=3

right=3: char='a'

	'a' in set! → remove s[left]='a', left++ → set={'b','c'}, left=1
	'a' not in set now → add 'a' → set={'b','c','a'}
	window = "bca", len=3, maxLen=3

right=4: char='b'

	'b' in set! → remove s[left]='b', left++ → set={'c','a'}, left=2
	'b' not in set now → add 'b' → set={'c','a','b'}
	window = "cab", len=3, maxLen=3

right=5: char='c'

	'c' in set! → remove s[left]='c', left++ → set={'a','b'}, left=3
	'c' not in set now → add 'c' → set={'a','b','c'}
	window = "abc", len=3, maxLen=3

right=6: char='b'

	'b' in set! → remove s[left]='a', left++ → set={'b','c'}, left=4
	'b' still in set! → remove s[left]='b', left++ → set={'c'}, left=5
	'b' not in set now → add 'b' → set={'c','b'}
	window = "cb", len=2, maxLen=3

right=7: char='b'

	'b' in set! → remove s[left]='c', left++ → set={'b'}, left=6
	'b' still in set! → remove s[left]='b', left++ → set={}, left=7
	'b' not in set now → add 'b' → set={'b'}
	window = "b", len=1, maxLen=3

Result: maxLen = 3 (substring "abc")
*/
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool)
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		// Shrink window from left while duplicate exists
		for charSet[s[right]] {
			delete(charSet, s[left])
			left++
		}

		// Add current character to set
		charSet[s[right]] = true

		// Update max length
		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}

func mainLongestSubstring() {
	test1 := "abcabcbb"
	fmt.Printf("Input: \"%s\" → Output: %d\n", test1, lengthOfLongestSubstring(test1)) // 3 ("abc")

	test2 := "bbbbb"
	fmt.Printf("Input: \"%s\" → Output: %d\n", test2, lengthOfLongestSubstring(test2)) // 1 ("b")

	test3 := "pwwkew"
	fmt.Printf("Input: \"%s\" → Output: %d\n", test3, lengthOfLongestSubstring(test3)) // 3 ("wke")

	test4 := ""
	fmt.Printf("Input: \"%s\" → Output: %d\n", test4, lengthOfLongestSubstring(test4)) // 0

	test5 := "dvdf"
	fmt.Printf("Input: \"%s\" → Output: %d\n", test5, lengthOfLongestSubstring(test5)) // 3 ("vdf")

	// Real-world example: Unique event sequence
	eventStream := "auth-login-login-logout-register"
	fmt.Printf("\nEvent stream: \"%s\"\n", eventStream)
	fmt.Printf("Longest unique sequence: %d characters\n", lengthOfLongestSubstring(eventStream))
}
