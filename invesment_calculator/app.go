package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64 = 1000
	expectedInterestRate := 5.5
	var years float64

	fmt.Print("Investment Amount:")
	fmt.Scan(&investmentAmount)

	fmt.Print("Investment Years:")
	fmt.Scan(&years)

	fmt.Print("Expected Interest Rate:")
	fmt.Scan(&expectedInterestRate)

	// FV = PV * (1 + r)^n
	value, actualValue := calculateFutureValues(investmentAmount, expectedInterestRate, years)

	formattedV := fmt.Sprintf("Future Value: %.2f\n", value)
	formattedAV := fmt.Sprintf("Future Value (adjusted with inflation): %.1f\n", actualValue)

	// fmt.Println(value)
	// fmt.Println(actualValue)
	// fmt.Printf("Future Value: %.2f\nFuture value(adjusted with inflation): %.1f\n", value, actualValue)
	fmt.Print(formattedV, formattedAV)
}

func calculateFutureValues(investmentAmount, expectedInterestRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedInterestRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
	// return
}