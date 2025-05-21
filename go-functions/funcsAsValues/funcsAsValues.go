package funcsAsValues

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// passing function as value to another function
	doubled := transformNumbers2(&numbers, double)
	fmt.Println("doubled: ", doubled)

	// passing function as value to another function
	tripled := transformNumbers2(&numbers, triple)
	fmt.Println("tripled: ", tripled)

	// Passing anonymous function as value to another function
	quadded := transformNumbers2(&numbers, func(n int) int {
		return n * 4
	})
	fmt.Println("quadded: ", quadded)
}

// passing function as value to another function
func transformNumbers2(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

// function returning another function as value
func getTransformFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
