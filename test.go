package main

import "fmt"

func myFunc(dividend int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("divisior is 0")
	}
	return dividend/divisor, nil
}

func main() {
	result, err := myFunc(10,0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
