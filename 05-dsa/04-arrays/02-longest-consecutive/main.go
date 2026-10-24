package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	hset := make(map[int]bool)
	for _, i := range nums {
		hset[i] = true
	}

	longest := 0
	for _, i := range nums {
		if hset[i-1] {
			continue
		}
		count := 1
		for hset[i+1] {
			count += 1
			i += 1
		}
		if count > longest {
			longest = count
		}
	}
	return longest
}

func main() {
	seq := []int{5}
	longest := longestConsecutive(seq)
	fmt.Println("Longest: ", longest)
}
