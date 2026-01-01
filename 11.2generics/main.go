package main

import "fmt"

// ==================SCENARIO 1 ===============================

/*
func printSlice(items []int) {
	for _, item := range items {
		fmt.Println(item)
		}
		}

func main(){
	printSlice([]int{1,2,3})
	printSlice([]string{"1,2,3"}) //Cant do this as printSlcie nly expects int, i would have to create new printslcie that accepts string
	}
*/

// ================== FIX ===============================

/*
Can dop any of the following, keeping T is convention
func printSlice[T any](items []T) {
func printSlice[T interface{}](items []T) {

Doing this is more secure, as we are restricting the types allowed
func printSlice[T int | string](items []T) {
func printSlice[T comparable](items []T) {
*/

func printSlice[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	printSlice([]int{1,2,3})
	printSlice([]string{"go","ts"})
	printSlice([]bool{true,false})
}