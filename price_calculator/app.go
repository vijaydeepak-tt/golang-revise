package main

import (
	"fmt"

	// "example.com/price_calculator/pkg/cmdmanager"
	"example.com/price_calculator/pkg/filemanager"
	"example.com/price_calculator/pkg/prices"
)

func main() {
	taxRates := []float64{0, 0.05, 0.1, 0.5}

	// results := make(map[float64][]float64)
	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		pricesJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := pricesJob.Process()

		if err != nil {
			fmt.Println("Could not process the job")
			fmt.Println(err)
		}
	}

}
