package main

import "fmt"

// 1. Regular function
func add(a int, b int) int {
	return a + b
}

// 2. Function returning multiple values
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return dividend / divisor, nil
}

// 3. Named return values (optional)
func split(num int) (half int, remainder int) {
	half = num / 2
	remainder = num % 2
	return // returns half, remainder
}

// 4. Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	fmt.Println(add(2, 3))
	fmt.Println(divide(10, 0))
	fmt.Println(split(7))

	// 5. Anonymous function, immediately invoked with argument 5
	result := func(x int) int {
		return x * x
	}(5)
	fmt.Println(result)

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
