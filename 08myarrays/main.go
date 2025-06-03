package main

import "fmt"

func main() {
	fmt.Println("Intro to arrays")
	// In Go, an array is a numbered sequence of elements of a specific length. 

	var fruitList [4]string // array of 4 strings
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Banana"

	fmt.Println("Fruit List:", fruitList) // prints the array, Fruit List: [Apple Orange  Banana], notice the empty string at index 2
	fmt.Println("Fruit List Length:", len(fruitList)) // prints the length of the array, Fruit List Length: 4, even though we have only 3 fruits in the array 

	var vegList = [3]string{"Potato", "Tomato", "Carrot"} // array of 3 strings
	fmt.Println("Veg List:", vegList) // prints the array, Veg List: [Potato Tomato Carrot]
	fmt.Println("Veg List Length:", len(vegList)) // prints the length of the array, Veg List Length: 3
}