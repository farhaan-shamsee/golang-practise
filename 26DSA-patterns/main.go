package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func printStar(n int) {
	for range n {
		for range n {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}

func printStar2(n int) {
	for i := 0; i < n; i++  {
		for j := 0; j <= i; j++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}

func printStar3(n int) {
	for i := 0; i < n; i++  {
		for j := 0; j <= i; j++ {
			fmt.Printf("%v ",j+1)
		}
		fmt.Printf("\n")
	}
}
func printStar4(n int) {
	for i := 0; i < n; i++  {
		for j := 0; j <= i; j++ {
			fmt.Printf("%v ",i+1) //we are just running the row number
		}
		fmt.Printf("\n")
	}
}

func printStar5(n int) {
	for i := 1; i <= n; i++  {
		for j := 0; j <n-i+1; j++ {
			fmt.Printf("* ") // formula n-row+1
		}
		fmt.Printf("\n")
	}
}

func printStar6(n int) {
	for i := 1; i <= n; i++  {
		for j := 1; j <=n-i+1; j++ {
			fmt.Printf("%v ",j) 
		}
		fmt.Printf("\n")
	}
}
func printStar7(n int) {
				// 	    *     
				//     ***    
				//    *****   
				//   *******  
				//  ********* 
				// ***********
				//  spaces,star,spaces , we count on each line how many space and how many star we want. and then create the formula for the star and spaces
	for i := 0; i < n; i++  {
		for j := 0; j <n-i-1; j++ {
			fmt.Printf(" ") 
		}
		for j := 0; j <2*i+1; j++ {
			fmt.Printf("*") 
		}
		for j := 0; j <n-i-1; j++ {
			fmt.Printf(" ") 
		}
		fmt.Printf("\n")
	}
}
func printStar8(n int) {
	for i := 0; i < n; i++  {
		for j := 0; j <i; j++ {
			fmt.Printf(" ") 
		}
		for j := 0; j <(2*n-(2*i+1)); j++ {
			fmt.Printf("*") 
		}
		for j := 0; j <i; j++ {
			fmt.Printf(" ") 
		}
		fmt.Printf("\n")
	}
}
func printStar9(n int) {
	for i := 0; i < n; i++  {
		for j := 0; j <i; j++ {
			fmt.Printf(" ") 
		}
		for j := 0; j <(2*n-(2*i+1)); j++ {
			fmt.Printf("*") 
		}
		for j := 0; j <i; j++ {
			fmt.Printf(" ") 
		}
		fmt.Printf("\n")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter number of lines you want to print:")
    input, _ := reader.ReadString('\n') // Reading by line
    input = strings.TrimSpace(input)     // Remove any surrounding whitespace

    number, err := strconv.Atoi(input) // Convert string input to an integer
    if err != nil {
        fmt.Println("Invalid input. Please enter a numeric value.")
        return
    }

    // printStar(number) // Pass the integer to printStar
    printStar8(number) // Pass the integer to printStar
}
