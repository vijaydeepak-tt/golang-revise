package main

import "fmt"

// by adding the * to the variable type it defines that argument is receiving the address of the value.
func add(a, b int, sum *int) {
	// to get the value to the address passed as argument, add * b4 the variable
	result := a + b

	// storing th result to the sum address
	// manipulating the data in the original sum variable
	*sum = result
}

func main() {
	a := 10
	b := 20
	sum := 0

	// creating the pointer variable for 'sum' variable
	sumPointer := &sum

	// passing the pointer variable of 'sum' as an argument
	add(a, b, sumPointer)

	fmt.Println("Sum:", sum)
}
