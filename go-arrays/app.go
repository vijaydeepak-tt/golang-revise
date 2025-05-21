package main

import "fmt"

// func main() {
// 	practiceMain()
// }

func main() {
	// userNames := []string{}
	// 2 is initial size, 5 is initial capacity. Wwe can add more than this sizes
	userNames := make([]string, 2, 5)

	userNames[0] = "Jones"

	userNames = append(userNames, "John")
	userNames = append(userNames, "Jane")
	// userNames = append(userNames, "Doe")
	// userNames = append(userNames, "Smith")
	// userNames = append(userNames, "Alice")
	// userNames = append(userNames, "Bob")
	// userNames = append(userNames, "Charlie")

	fmt.Println(userNames)

	for index, value := range userNames {
		fmt.Println(index, value)
	}
}

// func main() {
// 	prices := []float64{1.1, 2.2}
// 	fmt.Println(prices)

// 	// prices = append(prices, 3.4)
// 	updatedPrice := append(prices, 3.4)

// 	fmt.Println(prices, updatedPrice)

// 	discountedPrices := []float64{9.8, 8.7, 7.6, 6.5}

// 	updatedPrice = append(updatedPrice, discountedPrices...)

// 	fmt.Println(prices, updatedPrice)
// }

// func main() {
// 	// Arrays
// 	prices := [4]float64{1.1, 2.2, 3.3, 4.4}
// 	fmt.Println(prices)
// 	fmt.Println(prices[0])

// 	// slices
// 	futurePrices := prices[1:]
// 	fmt.Println(futurePrices)

// 	highestPrices := futurePrices[:1]
// 	fmt.Println(highestPrices)

// 	fmt.Println(len(highestPrices), cap(highestPrices)) // 1, 3

// 	highestPrices = prices[:1]
// 	fmt.Println(highestPrices)

// 	fmt.Println(len(highestPrices), cap(highestPrices)) // 1, 4
// }
