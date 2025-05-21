package closures

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	double := createTransformer(2)
	triple := createTransformer(3)

	doubled := transformNumbers(&numbers, double)
	fmt.Println("doubled: ", doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println("tripled: ", tripled)
}

// passing function as value to another function
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
