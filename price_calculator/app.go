package main

import (
	"fmt"

	// "example.com/price_calculator/pkg/cmdmanager"
	"example.com/price_calculator/pkg/filemanager"
	"example.com/price_calculator/pkg/prices"
)

func main() {
	taxRates := []float64{0, 0.05, 0.1, 0.5}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	// results := make(map[float64][]float64)
	for idx, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()

		pricesJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// err := pricesJob.Process()

		doneChans[idx] = make(chan bool)
		errorChans[idx] = make(chan error)
		// this change made the file gen in parallel, instead of 1 by 1
		go pricesJob.Process(doneChans[idx], errorChans[idx])

		// if err != nil {
		// 	fmt.Println("Could not process the job")
		// 	fmt.Println(err)
		// }
	}

	for index := range taxRates {
		// select case will wait for the first occurrence of any of the case channel, once occurred, then remaining will be skipped
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done with File no.", index+1)
		}
	}

}
