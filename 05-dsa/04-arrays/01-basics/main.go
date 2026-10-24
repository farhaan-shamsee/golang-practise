package main

import (
	"fmt"
)

func sumArr(arr []int) (int, int) {
	sum := 0
	product := 1
	for _, value := range arr {
		sum += value
		product *= value
	}
	return sum, product
}

func minMax(arr []int) (int, int) {
	min := arr[0]
	max := arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return min, max
}
func swapMinMax(arr []int) []int {
	min := arr[0]
	max := arr[0]
	minIndex := 0
	maxIndex := 0
	for index, val := range arr {
		if val < min {
			min = val
			minIndex = index
		}
		if val > max {
			max = val
			maxIndex = index
		}
	}
	arr[minIndex], arr[maxIndex] = arr[maxIndex], arr[minIndex]
	return arr
}

func uniq(arr []int) []int {
	result := []int{}
	for _, i := range arr {
		count := 0
		for _, j := range arr {
			if i == j {
				count += 1
			}
		}

		if count == 1 {
			result = append(result, i)
		}
	}
	return result
}

func intersection(arr1 []int, arr2 []int) []int {
	intArr := []int{}
	for _, val1 := range arr1 {
		foundInArr2 := false
		for _, val2 := range arr2 {
			if val1 == val2 {
				foundInArr2 = true
				break
			}
		}
		if foundInArr2 {
			alreadyAdded := false
			for _, existing := range intArr {
				if existing == val1 {
					alreadyAdded = true
					break
				}
			}
			if !alreadyAdded {
				intArr = append(intArr, val1)
			}
		}
	}
	return intArr
}

func singleNumber(nums []int) int {
	ans := 0
	for _, val := range nums {
		ans = ans ^ val
	}

	return ans
}

func main() {
	// sum, product := sumArr([]int{1,2,3})
	// fmt.Println("Sum: ", sum)
	// fmt.Println("Product: ", product)
	// min, max := minMax([]int{1, 2, -3, 40})
	// swappedArr := swapMinMax([]int{1, 200, -3, 40})
	// fmt.Println("Min: ", min)
	// fmt.Println("Max: ", max)
	// fmt.Println("Swapped: ", swappedArr)
	// uniq := uniq([]int{1,1,2,2,4})
	// fmt.Println("Uniq: ", uniq)
	interArr := intersection([]int{1, 1, 2, 2, 4}, []int{1, 1, 2, 2, 4, 5})
	fmt.Println("Intersection: ", interArr)
	single := singleNumber([]int{1, 1, 2, 2, 4})
	fmt.Println("Single Number: ", single)

}
