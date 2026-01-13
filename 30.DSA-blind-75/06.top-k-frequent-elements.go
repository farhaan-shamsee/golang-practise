/*
Problem: Top K Frequent Elements
Given an integer array nums and an integer k, return the k most frequent elements.
You may return the answer in any order.

Real-world use case for Platform Engineers:
- Finding top 5 most frequent error codes in log streams
- Identifying most resource-heavy pods in a cluster
- Detecting most common API endpoints being called
- Analyzing most frequent configuration patterns

Approaches:

1. Hash Map + Sorting:
  - Count frequencies using map
  - Sort by frequency
  - Take top k elements
  - Time: O(n log n), Space: O(n)

2. Hash Map + Heap (Min Heap):
  - Count frequencies using map
  - Maintain a min heap of size k
  - Time: O(n log k), Space: O(n + k)

3. Hash Map + Bucket Sort (Optimal):
  - Count frequencies using map
  - Create buckets where index = frequency
  - Collect top k from buckets (right to left)
  - Time: O(n), Space: O(n)

Algorithm (Bucket Sort Approach):
1. Count frequency of each element using hash map
2. Create array of buckets where buckets[i] contains all elements with frequency i
3. Iterate buckets from right to left (highest frequency first)
4. Collect elements until we have k elements

Time Complexity: O(n) - count frequencies + bucket sort
Space Complexity: O(n) - frequency map + buckets array

Example: nums = [1,1,1,2,2,3], k = 2

Step 1 - Count frequencies:

	freq = {1: 3, 2: 2, 3: 1}

Step 2 - Create buckets (index = frequency):

	buckets[0] = []
	buckets[1] = [3]        // elements with freq 1
	buckets[2] = [2]        // elements with freq 2
	buckets[3] = [1]        // elements with freq 3
	buckets[4] = []
	buckets[5] = []
	buckets[6] = []

Step 3 - Collect top k from right to left:

	Start from highest frequency (6 down to 0)
	buckets[6] = [] → skip
	buckets[5] = [] → skip
	buckets[4] = [] → skip
	buckets[3] = [1] → add 1, result = [1], count = 1
	buckets[2] = [2] → add 2, result = [1, 2], count = 2 (reached k)

Result: [1, 2] (order doesn't matter, could be [2, 1])

Verification:

	1 appears 3 times ✓ (most frequent)
	2 appears 2 times ✓ (second most frequent)
*/
package main

import (
	"fmt"
	"sort"
)

// Approach 1: Hash Map + Sorting
// Time: O(n log n), Space: O(n)
func topKFrequentSort(nums []int, k int) []int {
	// Step 1: Count frequencies
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// Step 2: Create slice of [number, frequency] pairs
	type pair struct {
		num   int
		count int
	}
	pairs := []pair{}
	for num, count := range freq {
		pairs = append(pairs, pair{num, count})
	}

	// Step 3: Sort by frequency (descending)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	// Step 4: Take top k elements
	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, pairs[i].num)
	}

	return result
}

// Approach 2: Hash Map + Min Heap
// Time: O(n log k), Space: O(n + k)
type heapItem struct {
	num   int
	count int
}

type minHeap []heapItem

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(heapItem))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func topKFrequentHeap(nums []int, k int) []int {
	// Step 1: Count frequencies
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// Step 2: Use min heap of size k
	h := &minHeap{}

	for num, count := range freq {
		if h.Len() < k {
			// Heap not full, add element
			*h = append(*h, heapItem{num, count})
			if h.Len() == k {
				// Just filled, heapify
				for i := k/2 - 1; i >= 0; i-- {
					heapifyDown(h, i)
				}
			}
		} else if count > (*h)[0].count {
			// Current element more frequent than min in heap
			(*h)[0] = heapItem{num, count}
			heapifyDown(h, 0)
		}
	}

	// Step 3: Extract elements from heap
	result := []int{}
	for _, item := range *h {
		result = append(result, item.num)
	}

	return result
}

func heapifyDown(h *minHeap, i int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < h.Len() && (*h)[left].count < (*h)[smallest].count {
		smallest = left
	}
	if right < h.Len() && (*h)[right].count < (*h)[smallest].count {
		smallest = right
	}

	if smallest != i {
		(*h)[i], (*h)[smallest] = (*h)[smallest], (*h)[i]
		heapifyDown(h, smallest)
	}
}

// Approach 3: Hash Map + Bucket Sort (Optimal)
// Time: O(n), Space: O(n)
func topKFrequent(nums []int, k int) []int {
	// Step 1: Count frequencies
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// Step 2: Create buckets (index = frequency)
	// buckets[i] contains all numbers that appear i times
	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	// Step 3: Collect top k elements from highest frequency
	result := []int{}
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		result = append(result, buckets[i]...)
	}

	return result[:k]
}

func mainTopKFrequent() {
	nums1 := []int{1, 1, 1, 2, 2, 3}
	k1 := 2
	fmt.Printf("Input: nums = %v, k = %d\n", nums1, k1)
	fmt.Printf("Approach 1 (Sort):   %v\n", topKFrequentSort(nums1, k1))
	fmt.Printf("Approach 2 (Heap):   %v\n", topKFrequentHeap(nums1, k1))
	fmt.Printf("Approach 3 (Bucket): %v\n\n", topKFrequent(nums1, k1))

	nums2 := []int{1}
	k2 := 1
	fmt.Printf("Input: nums = %v, k = %d\n", nums2, k2)
	fmt.Printf("Approach 1 (Sort):   %v\n", topKFrequentSort(nums2, k2))
	fmt.Printf("Approach 2 (Heap):   %v\n", topKFrequentHeap(nums2, k2))
	fmt.Printf("Approach 3 (Bucket): %v\n\n", topKFrequent(nums2, k2))

	// Real-world example: Top 3 error codes
	errorCodes := []int{404, 500, 404, 404, 200, 500, 403, 404, 500, 500}
	k3 := 3
	fmt.Printf("Input: error codes = %v, k = %d\n", errorCodes, k3)
	fmt.Printf("Approach 1 (Sort):   %v\n", topKFrequentSort(errorCodes, k3))
	fmt.Printf("Approach 2 (Heap):   %v\n", topKFrequentHeap(errorCodes, k3))
	fmt.Printf("Approach 3 (Bucket): %v\n", topKFrequent(errorCodes, k3))

	fmt.Println("\n--- Complexity Comparison ---")
	fmt.Println("Approach 1 - Sort:   Time O(n log n), Space O(n)")
	fmt.Println("Approach 2 - Heap:   Time O(n log k), Space O(n + k)")
	fmt.Println("Approach 3 - Bucket: Time O(n), Space O(n) ← Best!")
}
