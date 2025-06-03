package main

import "fmt"

func main() {
	fmt.Println("Intro to pointers")
	
	myNumber := 23


	var ptr *int // ptr is a pointer to an int. This is the first usage of *, to declare a variable as pointer.
	// fmt.Println("Value of pointer is:", ptr)


	ptr = &myNumber // ptr is a pointer to myNumber , referencing the address of myNumber that is why we use & operator.
	// here ptr now has the memory address of myNumber.
	fmt.Println("Value of pointer is:", ptr) // prints the address of myNumber
	fmt.Println("Value of myNumber is:", *ptr) // dereferencing the pointer to get the value of myNumber.
	// This is the second use of the *. If its in front of a variable(in 1st case it was in front of a datatype) then it shows the 
	// value stored at that memory address. This process of accessing the value through the pointer is called dereferencing.

	*ptr = *ptr + 2 // dereferencing the pointer to change the value of myNumber. This is the third usage of *.
	fmt.Println("Value of myNumber is:", myNumber) // prints 25
}



