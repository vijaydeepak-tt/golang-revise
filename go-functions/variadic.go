package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Sum of numbers:", sum)

	anotherSum := sumAll(numbers...)
	fmt.Println("Sum of numbers:", anotherSum)

	multiplied := multipleAll(2, 1, 2, 3, 4, 5)
	fmt.Println("Multiplied numbers:", multiplied)
}

func sumAll(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func multipleAll(firstValue int, numbers ...int) int {
	multiplied := firstValue

	for _, val := range numbers {
		multiplied *= val
	}

	return multiplied
}
