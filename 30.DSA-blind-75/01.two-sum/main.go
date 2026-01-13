/*
brute force: compare 1st element with next element and check if it is sum
optimal: hashMap: take 1st element, find complement = target-1st element, check if comp exist in map, if yes then return otherwise store 1st number.
*/
package main

import "fmt"

func twoSumBrute(nums []int, target int) []int {
	// Brute Force Approach: Check every possible pair
	fmt.Println("\n--- Brute Force Approach ---")
	n := len(nums)

	// Outer loop: pick first element
	for i := 0; i < n-1; i++ {
		fmt.Printf("Checking nums[%d] = %d\n", i, nums[i])

		// Inner loop: pick second element (after first)
		for j := i + 1; j < n; j++ {
			sum := nums[i] + nums[j]
			fmt.Printf("  nums[%d] + nums[%d] = %d + %d = %d", i, j, nums[i], nums[j], sum)

			if sum == target {
				fmt.Printf(" ✓ FOUND!\n")
				return []int{i, j}
			}
			fmt.Printf(" (not target)\n")
		}
	}
	fmt.Println("No pair found")
	return []int{}
}

func twoSumOptimal(nums []int, target int) []int {
	// Optimal Approach: Use hashmap to find complement in O(1)
	fmt.Println("\n--- Optimal Approach (HashMap) ---")

	// Create a map to store: value -> index
	m := make(map[int]int)
	fmt.Println("HashMap initialized: {}")

	// Iterate through array once
	for i, num := range nums {
		// Calculate what number we need to reach target
		comp := target - num
		// fmt.Printf("\nStep %d: nums[%d] = %d\n", i+1, i, num)
		// fmt.Printf("  Need complement: %d - %d = %d\n", target, num, comp)

		// Check if complement exists in map
		if j, ok := m[comp]; ok {
			fmt.Printf("  ✓ Found complement %d at index %d in map!\n", comp, j)
			fmt.Printf("  Answer: [%d, %d]\n", j, i)
			return []int{j, i}
		}

		// Store current number and its index in map
		fmt.Printf("  Complement not found, storing: map[%d] = %d\n", num, i)
		m[num] = i
		fmt.Printf("  HashMap now: %v\n", m)
	}

	fmt.Println("No pair found")
	return []int{}
}

func main() {
	// Example 1
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println("========================================")
	fmt.Printf("Example 1: nums = %v, target = %d\n", nums, target)
	fmt.Println("========================================")
	result1 := twoSumBrute(nums, target)
	fmt.Printf("\nResult: %v\n", result1)

	result2 := twoSumOptimal(nums, target)
	fmt.Printf("\nResult: %v\n", result2)

	// Example 2
	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Println("\n========================================")
	fmt.Printf("Example 2: nums = %v, target = %d\n", nums2, target2)
	fmt.Println("========================================")
	result3 := twoSumBrute(nums2, target2)
	fmt.Printf("\nResult: %v\n", result3)

	result4 := twoSumOptimal(nums2, target2)
	fmt.Printf("\nResult: %v\n", result4)
}
