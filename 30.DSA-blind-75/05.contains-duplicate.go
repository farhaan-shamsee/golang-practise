/*
Problem: Contains Duplicate
Given an integer array nums, return true if any value appears at least twice
in the array, and return false if every element is distinct.

Approaches:

1. Brute Force (Nested Loops):
  - Compare each element with every other element
  - Time: O(n²), Space: O(1)

2. Sorting:
  - Sort the array first
  - Check adjacent elements for duplicates
  - Time: O(n log n), Space: O(1) or O(n) depending on sort implementation

3. Hash Set (Optimal):
  - Use a map/set to track seen elements
  - For each element, check if it exists in set
  - If yes, return true (duplicate found)
  - If no, add to set and continue
  - Time: O(n), Space: O(n)

Algorithm (Hash Set Approach):
1. Create an empty map/set
2. Loop through each number in array
3. Check if number exists in map
  - If yes: return true (duplicate found)
  - If no: add number to map

4. If loop completes: return false (no duplicates)

Time Complexity: O(n) - single pass through array
Space Complexity: O(n) - worst case, all elements are unique

Example 1: nums = [1,2,3,1]

	i=0: num=1, seen={}, add 1 → seen={1}
	i=1: num=2, seen={1}, add 2 → seen={1,2}
	i=2: num=3, seen={1,2}, add 3 → seen={1,2,3}
	i=3: num=1, seen={1,2,3}, 1 exists! → return true

Example 2: nums = [1,2,3,4]

	i=0: num=1, seen={}, add 1 → seen={1}
	i=1: num=2, seen={1}, add 2 → seen={1,2}
	i=2: num=3, seen={1,2}, add 3 → seen={1,2,3}
	i=3: num=4, seen={1,2,3}, add 4 → seen={1,2,3,4}
	Loop ends → return false (all distinct)
*/
package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true // Duplicate found
		}
		seen[num] = true
	}

	return false // No duplicates
}

func mainCall() {
	nums1 := []int{1, 2, 3, 1}
	fmt.Printf("Input: %v\n", nums1)
	fmt.Printf("Output: %v\n\n", containsDuplicate(nums1)) // true

	nums2 := []int{1, 2, 3, 4}
	fmt.Printf("Input: %v\n", nums2)
	fmt.Printf("Output: %v\n\n", containsDuplicate(nums2)) // false

	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Printf("Input: %v\n", nums3)
	fmt.Printf("Output: %v\n", containsDuplicate(nums3)) // true
}
