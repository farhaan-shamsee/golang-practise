package main

import "fmt"

// myRec prints "farrow" to the standard output for each integer value from i up to n (inclusive).
//
// It uses recursion to increment the value of i until it exceeds n, at which point the recursion terminates.
//
// Parameters:
//   - i: The starting integer value.
//   - n: The ending integer value.
//
// Notes on Recursion:
//   - The function calls itself with an incremented value of i (i+1) until the base case (i > n) is met.
//   - Each recursive call prints "farrow" before making the next call.
//   - Recursion is a technique where a function calls itself to solve a problem by breaking it down into smaller subproblems.
//   - Always ensure a base case (here, i > n) to prevent infinite recursion and potential stack overflow.
//
// Time Complexity: O(n - i + 1), where n and i are the input parameters. Each call prints once and recurses until i > n.
// Space Complexity: O(n - i + 1), due to the call stack created by recursion.
func myRec(i int, n int) {
	if i > n {
		return
	}
	fmt.Println("farrow")
	myRec(i+1, n)
}

func revNum(i int, j int) {
	if i < 1 {
		return
	}
	fmt.Println(i)
	revNum(i-1, j)
}

func revNumBacktrack(i int, j int) {
	if i > j {
		return
	}
	revNumBacktrack(i+1, j)
	fmt.Println(i)
}

func sumOfN(i int, sum int) {
	if i < 1 {
		fmt.Println(sum)
		return
	}
	sumOfN(i-1, sum+i)

}

func sumOfNFunctional(i int) int {
	if i == 0 {
		return 0
	}
	return i + sumOfNFunctional(i-1)
}

func factorial(i int) int {
	if i == 0 {
		return 1 // see we are returning 1, if it were 0 then at the end we would have returned 1*0 which would make everything 0.
	}
	return i * factorial(i-1)
}

func main() {
	n := 3
	fmt.Println("test")
	// myRec(1, n)
	// revNum(n, n)
	// revNumBacktrack(1, n)
	sumOfN(n, 0)
	fmt.Println(sumOfNFunctional(n))
	fmt.Println(factorial(n))
}
