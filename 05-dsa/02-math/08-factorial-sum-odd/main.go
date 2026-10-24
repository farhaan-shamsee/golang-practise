package main

import (
	"fmt"
)

func primeOrNot(n int) {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			fmt.Printf("%d is not prime\n", n)
			return
		}
	}
	fmt.Printf("%d is prime\n", n)
}

func factorial(n int) {
	fact := 1
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	fmt.Println(fact)
}

func sumOdd(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		i++
	}
	fmt.Println(sum)
}

func main() {
	fmt.Println("Hello, World!")
	// primeOrNot(10)
	// factorial(3)
	sumOdd(5)
}
