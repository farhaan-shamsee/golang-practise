package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printStar(n int) {
	// Prints:
	// * * * * *
	// * * * * *
	// * * * * *
	// * * * * *
	// * * * * *
	for range n {
		for range n {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}

func printStar2(n int) {
	// Prints increasing rows:
	// *
	// * *
	// * * *
	// * * * *
	// * * * * *
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}

func printStar3(n int) {
	// Prints row numbers in a growing triangle:
	// 1
	// 1 2
	// 1 2 3
	// 1 2 3 4
	// 1 2 3 4 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%v ", j+1)
		}
		fmt.Printf("\n")
	}
}

func printStar4(n int) {
	// Prints the same row number repeated on each row:
	// 1
	// 2 2
	// 3 3 3
	// 4 4 4 4
	// 5 5 5 5 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%v ", i+1)
		}
		fmt.Printf("\n")
	}
}

func printStar5(n int) {
	// Prints decreasing stars from left to right:
	// * * * * *
	// * * * *
	// * * *
	// * *
	// *
	for i := 1; i <= n; i++ {
		for j := 0; j < n-i+1; j++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}

func printStar6(n int) {
	// Prints descending number rows:
	// 1 2 3 4 5
	// 1 2 3 4
	// 1 2 3
	// 1 2
	// 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i+1; j++ {
			fmt.Printf("%v ", j)
		}
		fmt.Printf("\n")
	}
}

func printStar7(n int) {
	// Prints a center-aligned pyramid:
	//     *
	//    ***
	//   *****
	//  *******
	// *********
	// spaces decrease, stars increase by 2 each row
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func printStar8(n int) {
	// Prints an inverted pyramid:
	// *********
	//  *******
	//   *****
	//    ***
	//     *
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < (2*n - (2*i + 1)); j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func printStar9(n int) {
	// Prints a mountain shape: grows up, then shrinks down.
	// *
	// **
	// ***
	// ****
	// *****
	// ****
	// ***
	// **
	// *
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	for i := n - 1; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func printPattern01(n int) {
	// Prints alternating 1 and 0 like a checkerboard triangle:
	// 1
	// 0 1
	// 1 0 1
	// 0 1 0 1
	// 1 0 1 0 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if (i+j)%2 == 0 {
				fmt.Printf("1 ")
			} else {
				fmt.Printf("0 ")
			}
		}
		fmt.Printf("\n")
	}
}

func printPatternMirrorNumber(n int) {
	// Prints mirrored increasing numbers:
	// 1        1
	// 12      21
	// 123    321
	// 1234  4321
	// 1234554321
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d", j)
		}
		for j := 0; j < 2*(n-i); j++ {
			fmt.Printf(" ")
		}
		for j := i; j >= 1; j-- {
			fmt.Printf("%d", j)
		}
		fmt.Printf("\n")
	}
}

func printPatternSerial(n int) {
	// Prints values in row-wise increasing order:
	// 1
	// 2 3
	// 4 5 6
	// 7 8 9 10
	// 11 12 13 14 15
	value := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", value)
			value++
		}
		fmt.Printf("\n")
	}
}

func printPatternAlphabetIncreasing(n int) {
	// Pattern:
	// A
	// A B
	// A B C
	// A B C D
	// A B C D E
	for i := 0; i < n; i++ {
		for ch := 'A'; ch < 'A'+rune(i+1); ch++ {
			fmt.Printf("%c ", ch)
		}
		fmt.Printf("\n")
	}
}

func printPatternAlphabetDecreasing(n int) {
	// Pattern:
	// A B C D E
	// A B C D
	// A B C
	// A B
	// A
	for i := 0; i < n; i++ {
		for ch := 'A'; ch < 'A'+rune(n-i); ch++ {
			fmt.Printf("%c ", ch)
		}
		fmt.Printf("\n")
	}
}

func patternAlphabetSequence(n int) {
	// Pattern:
	// A
	// B B
	// C C C
	// D D D D
	// E E E E E
	char := 'A'
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%c ", char)
		}
		fmt.Printf("\n")
		char++
	}
}

func patternAlphabetMirror(n int) {
	// Pattern (n=5) — char resets to 'A' every row, rises to a peak, then mirrors back down:
	//       A
	//      A B A
	//     A B C B A
	//    A B C D C B A
	//   A B C D E D C B A
	for i := 0; i < n; i++ {
		for j := 0; j <= n-i; j++ {
			fmt.Printf(" ")
		}
		char := 'A'
		breakpoint := (2*i + 1) / 2
		for j := 1; j <= 2*i+1; j++ {
			fmt.Printf("%c ", char)
			if j <= breakpoint {
				char++
			} else {
				char--
			}
		}
		fmt.Printf("\n")
	}
}

func pattern(n int) {
	for i := 0; i < n; i++ {
		for char := rune('E' - i); char <= 'E'; char++ {
			fmt.Printf("%c ", char)
		}
		fmt.Printf("\n")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number of lines you want to print:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a numeric value.")
		return
	}

	// printStar(number)
	printPatternSerial(number)
}
