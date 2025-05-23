package main

import "fmt"

func main() {
	fmt.Println("Intro to pointers")

	// var ptr *int // ptr is a pointer to an int
	// fmt.Println("Value of pointer is:", ptr)

	myNumber := 23

	var ptr = &myNumber // ptr is a pointer to myNumber , referencing the address of myNumber that is why we use & operator
	fmt.Println("Value of pointer is:", ptr) // prints the address of myNumber
	fmt.Println("Value of myNumber is:", *ptr) // dereferencing the pointer to get the value of myNumber

	*ptr = *ptr + 2 // dereferencing the pointer to change the value of myNumber
	fmt.Println("Value of myNumber is:", myNumber) // prints 25
}



