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
    printStar2(number) // Pass the integer to printStar
}
