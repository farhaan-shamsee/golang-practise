/*
Problem: Group Anagrams
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
using all the original letters exactly once.

Real-world use case for Platform Engineers:
- Grouping similar metadata or configurations
- Clustering pods with similar labels/annotations
- Organizing services with similar characteristics
- Grouping log patterns with similar keywords

Approaches:

1. Sort Each String (Common):
   - For each string, sort its characters
   - Use sorted string as key in hash map
   - Group strings with same sorted key
   - Time: O(n * k log k) where n = number of strings, k = max length
   - Space: O(n * k)

2. Character Count as Key (Optimal):
   - Count character frequency for each string
   - Use frequency array/map as key
   - Time: O(n * k) where k = average string length
   - Space: O(n * k)

Algorithm (Sorting Approach):
1. Create a hash map where key = sorted string, value = list of original strings
2. For each string in input:
   a. Sort the characters
   b. Use sorted version as key
   c. Append original string to that key's list
3. Return all values from the map

Time Complexity: O(n * k log k) where n = strings count, k = max string length
Space Complexity: O(n * k) for storing all strings

Example: strs = ["eat","tea","tan","ate","nat","bat"]

Step 1 - Process each string:
  "eat" → sort → "aet" → map["aet"] = ["eat"]
  "tea" → sort → "aet" → map["aet"] = ["eat", "tea"]
  "tan" → sort → "ant" → map["ant"] = ["tan"]
  "ate" → sort → "aet" → map["aet"] = ["eat", "tea", "ate"]
  "nat" → sort → "ant" → map["ant"] = ["tan", "nat"]
  "bat" → sort → "abt" → map["abt"] = ["bat"]

Step 2 - Final map state:
  map = {
    "aet": ["eat", "tea", "ate"],
    "ant": ["tan", "nat"],
    "abt": ["bat"]
  }

Step 3 - Extract values:
  Result: [["eat","tea","ate"], ["tan","nat"], ["bat"]]

Verification:
  "eat", "tea", "ate" are anagrams (same letters) ✓
  "tan", "nat" are anagrams ✓
  "bat" stands alone ✓
*/
package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	// Map: sorted string → list of anagrams
	groups := make(map[string][]string)

	for _, str := range strs {
		// Sort the string to use as key
		sortedStr := sortString(str)
		groups[sortedStr] = append(groups[sortedStr], str)
	}

	// Extract all groups
	result := [][]string{}
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// Helper function to sort a string
func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func mainGroupAnagram() {
	strs1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("Input: %v\n", strs1)
	fmt.Printf("Output: %v\n\n", groupAnagrams(strs1))
	// [["eat","tea","ate"], ["tan","nat"], ["bat"]]

	strs2 := []string{""}
	fmt.Printf("Input: %v\n", strs2)
	fmt.Printf("Output: %v\n\n", groupAnagrams(strs2))
	// [[""]]

	strs3 := []string{"a"}
	fmt.Printf("Input: %v\n", strs3)
	fmt.Printf("Output: %v\n\n", groupAnagrams(strs3))
	// [["a"]]

	// Real-world example: Group similar pod labels
	podLabels := []string{"app-prod", "prod-app", "db-staging", "staging-db", "cache"}
	fmt.Printf("Pod labels: %v\n", podLabels)
	fmt.Printf("Grouped: %v\n", groupAnagrams(podLabels))
	// Groups labels with same characters together
}
