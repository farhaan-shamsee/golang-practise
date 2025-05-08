package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Intro to slices")

	var fruitList = []string{"Apple", "Tomato", "Peach" }
	fmt.Printf("Type of fruitlist is %T\n", fruitList) // slice of strings

	fruitList = append(fruitList, "Banana", "Mango") // append a new fruit to the slice
	fmt.Println("Fruit List:", fruitList) // prints the slice, Fruit List: [Apple Tomato Peach Banana]
	fmt.Println("Fruit List Length:", len(fruitList)) // prints the length of the slice, Fruit List Length: 4

	// fruitList = append(fruitList[1:]) // remove the first fruit from the slice
	// fmt.Println("Fruit List:", fruitList) // prints the slice, Fruit List: [Tomato Peach Banana]
	// fruitList = append(fruitList[:2]) // remove the last fruit from the slice
	// fmt.Println("Fruit List:", fruitList) // prints the slice, Fruit List: [Tomato Peach]

	highScores := make([]int, 4) // create a slice of 4 integers

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 // this will cause a runtime error: index out of range

	highScores = append(highScores, 777, 666, 321) // append a new score to the slice. All of the memory is reallocated to accommodate the new scores
	fmt.Println("High Scores:", highScores) // prints the slice, High Scores: [234 945 465 867]

	sort.Ints(highScores) // sort the slice in ascending order
	fmt.Println("High Scores Sorted:", highScores) // prints the slice, High Scores Sorted: [234 321 465 666 777 867]
	fmt.Println(sort.IntsAreSorted(highScores)) // prints true if the slice is sorted, false otherwise

}