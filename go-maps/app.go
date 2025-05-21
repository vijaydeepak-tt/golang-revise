package main

import "fmt"

type flatMap map[string]float64

func (fm flatMap) output() {
	fmt.Println("FlatMap output:", fm)
}

func main() {
	// courseRating := map[string]float64{}
	// courseRating := make(map[string]float64, 3)
	courseRating := make(flatMap, 3)

	courseRating["Go"] = 4.5
	courseRating["Python"] = 4.7
	courseRating["Java"] = 4.2

	// fmt.Println(courseRating)
	courseRating.output()

	for key, value := range courseRating {
		fmt.Printf("Key: %s, Value: %.2f\n", key, value)
	}
}

// func main() {
// 	websites := map[string]string{
// 		"Google":   "https://www.google.com",
// 		"Facebook": "https://www.facebook.com",
// 		"Twitter":  "https://www.twitter.com",
// 	}

// 	fmt.Println(websites)
// 	fmt.Println(websites["Google"])

// 	websites["Instagram"] = "https://www.instagram.com"
// 	fmt.Println(websites)

// 	// Delete a key in map
// 	delete(websites, "Facebook")
// 	fmt.Println(websites)

// 	// check if the key exists
// 	if _, ok := websites["Facebook"]; ok {
// 		fmt.Println("Facebook exists in the map")
// 	} else {
// 		fmt.Println("Facebook does not exist in the map")
// 	}

// }
