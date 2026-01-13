/*
Problem: Product of Array Except Self
Given an integer array nums, return an array answer such that answer[i] is equal to
the product of all elements of nums except nums[i].

Constraint: Must solve in O(n) time without using division.

Approach: Two-Pass Algorithm
Instead of calculating the product of all elements and dividing by nums[i],
we use left and right products.

For each position i:
- Left product = product of all elements before i (nums[0] * nums[1] * ... * nums[i-1])
- Right product = product of all elements after i (nums[i+1] * nums[i+2] * ... * nums[n-1])
- answer[i] = leftProduct[i] * rightProduct[i]

Algorithm:
1. Create answer array of same size
2. Pass 1 (Left to Right): Build left products
  - answer[0] = 1 (nothing to the left)
  - answer[i] = answer[i-1] * nums[i-1] (product of all elements before i)

3. Pass 2 (Right to Left): Multiply by right products
  - Start with right = 1 (nothing to the right of last element)
  - answer[i] *= right (multiply existing left product by right product)
  - Update right = right * nums[i] (for next position)

Time Complexity: O(n) - two passes through array
Space Complexity: O(1) - only output array (doesn't count as extra space)

Example: nums = [1,2,3,4]
Pass 1 (Left products):

	i=0: answer[0] = 1                    → [1, _, _, _]
	i=1: answer[1] = 1 * nums[0] = 1      → [1, 1, _, _]
	i=2: answer[2] = 1 * nums[1] = 2      → [1, 1, 2, _]
	i=3: answer[3] = 2 * nums[2] = 6      → [1, 1, 2, 6]

Pass 2 (Right products):

	right=1
	i=3: answer[3] = 6 * 1 = 6,   right = 1 * 4 = 4   → [1, 1, 2, 6]
	i=2: answer[2] = 2 * 4 = 8,   right = 4 * 3 = 12  → [1, 1, 8, 6]
	i=1: answer[1] = 1 * 12 = 12, right = 12 * 2 = 24 → [1, 12, 8, 6]
	i=0: answer[0] = 1 * 24 = 24, right = 24 * 1 = 24 → [24, 12, 8, 6]

Result: [24, 12, 8, 6]
Verification:

	answer[0] = 2*3*4 = 24 ✓
	answer[1] = 1*3*4 = 12 ✓
	answer[2] = 1*2*4 = 8  ✓
	answer[3] = 1*2*3 = 6  ✓
*/
package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// Pass 1: Left products (everything before i)
	answer[0] = 1 // Nothing left of 0
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	// Pass 2: Multiply by right products (everything after i)
	right := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= right // left × right
		right *= nums[i]   // for next position
	}

	return answer
}

func runProductExceptSelf() {
	nums := []int{1, 2, 3, 4}
	fmt.Printf("Input: %v\n", nums)
	fmt.Printf("Output: %v\n", productExceptSelf(nums)) // [24, 12, 8, 6]

	nums2 := []int{-1, 1, 0, -3, 3}
	fmt.Printf("\nInput: %v\n", nums2)
	fmt.Printf("Output: %v\n", productExceptSelf(nums2)) // [0, 0, 9, 0, 0]
}
