package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func practiceMain() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]string{"Chess", "Books", "Programming"}
	fmt.Println("1", hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println("2.1", hobbies[0])
	// fmt.Println("2.2", hobbies[1:3])
	fmt.Println("2.2", hobbies[1:])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	hobbiesSlice := hobbies[0:2]
	// hobbiesSlice := hobbies[:2]
	fmt.Println("3", hobbiesSlice)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	hobbiesSlice = hobbiesSlice[1:3]
	// hobbiesSlice = hobbiesSlice[1:]
	fmt.Println("4", hobbiesSlice)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"FIRE", "NOTHING"}

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "TRAVEL"
	goals = append(goals, "FAMILY")

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	products := []Product{{title: "One", id: 1, price: 1.1}, {title: "Two", id: 2, price: 2.1}}
	products = append(products, Product{title: "Three", id: 3, price: 3.2})

	fmt.Println(products)
}
