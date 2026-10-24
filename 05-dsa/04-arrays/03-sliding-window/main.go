package main

import (
	"fmt"
)

func slidingWindow(arr []int, winSize int) (int, error) {
	if winSize <= 0 || winSize > len(arr) {
		return 0, fmt.Errorf("Invalid window size")
	}
	oldSum := 0
	for index, val := range arr {
		if index < winSize {
			oldSum = oldSum + val
		}
	}
	newSum := 0
	max := oldSum

	for i := winSize; i < len(arr); i++ {
		newSum = oldSum - arr[i-winSize] + arr[i]
		if newSum > max {
			max = newSum
		}

		oldSum = newSum
	}
	return max, nil
}
func main() {
	windowSize := 6
	result, err := slidingWindow([]int{4, 2, 7, 1, 8}, windowSize)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Println("Result: ", result)
}
